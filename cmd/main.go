package main

import (
	"context"
	"fmt"
	"genproto/inventory_service"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Invan2/invan_inventory_service/config"
	"github.com/Invan2/invan_inventory_service/events"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/services/listeners"
	"github.com/Invan2/invan_inventory_service/storage"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"google.golang.org/grpc"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	cfg := config.Load()
	// log := logger.New(cfg.LogLevel, cfg.ServiceName)
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)

	log.Info("config", logger.Any("data", cfg), logger.Any("env", os.Environ()))

	defer logger.Cleanup(log)
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	eClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: cfg.ElasticSearchUrls,
		Username:  cfg.ElasticSearchUser,
		Password:  cfg.ElasticSearchPassword,
	})
	if err != nil {
		log.Error("elastic", logger.Error(err))
		return
	}

	_, err = eClient.Ping()
	if err != nil {
		log.Error("elastic ping", logger.Error(err))
		return
	}
	postgresURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	conn, err := sqlx.Connect("postgres", postgresURL)
	if err != nil {
		log.Error("error while connecting to postgres", logger.Error(err))
		return
	}
	defer conn.Close()

	elastic := storage.NewStorageES(log, eClient)

	strg, err := storage.NewStoragePg(ctx, log, conn)
	if err != nil {
		log.Error("error while create storage", logger.Error(err))
		return
	}

	conf := kafka.ConfigMap{
		"bootstrap.servers":                     cfg.KafkaUrl,
		"group.id":                              config.ConsumerGroupID,
		"auto.offset.reset":                     "earliest",
		"go.events.channel.size":                1000000,
		"socket.keepalive.enable":               true,
		"metadata.max.age.ms":                   86400000,
		"metadata.request.timeout.ms":           30000,
		"retries":                               1000000,
		"message.timeout.ms":                    300000,
		"socket.timeout.ms":                     30000,
		"max.in.flight.requests.per.connection": 5,
		"heartbeat.interval.ms":                 3000,
		"enable.idempotence":                    true,
		"message.max.bytes":                     1000000000,
	}

	log.Info("kafka config", logger.Any("config", conf))

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Error("error while creating producer", logger.Error(err))
		return
	}

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Error("error while creating consumer", logger.Error(err))
		return
	}

	pubsubServer, err := events.NewPubSubServer(log, producer, consumer, strg, elastic)
	if err != nil {
		log.Fatal("error creating pubSubServer", logger.Error(err))
		return
	}

	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*20),
		grpc.MaxSendMsgSize(1024*1024*20),
	)

	inventory_service.RegisterInventoryServiceServer(server, listeners.NewInventoryService(strg, log, pubsubServer, elastic))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", cfg.HttpHost, cfg.HttpPort))
	if err != nil {
		log.Error("http", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
		if err := pubsubServer.Shutdown(); err != nil {
			log.Error("error while shutdown pub sub server")
			return
		}
	}()

	go func() {
		if err := pubsubServer.Run(ctx); err != nil {
			log.Error("error while start pub sub server", logger.Error(err))
			panic(err)
		}
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatal("serve", logger.Error(err))
		return
	}

}
