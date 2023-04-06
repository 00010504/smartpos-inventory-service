package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) UpsertProduct(ctx context.Context, event *kafka.Message) error {

	var req common.CreateProductCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create product event", logger.Any("event", req))

	tr, err := e.strg.WithTransaction(ctx)
	if err != nil {
		return errors.Wrap(err, "error while run transaction")
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	err = tr.Product().Upsert(ctx, &req)
	if err != nil {
		e.log.Error("error while upsert product", logger.Error(err))
		return err
	}

	err = tr.Import().Upsert(ctx, &req)
	if err != nil {
		e.log.Error("error while upsert import", logger.Error(err))
		return err
	}

	return nil
}
