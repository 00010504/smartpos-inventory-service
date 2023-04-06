package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) CreateCompany(ctx context.Context, event *kafka.Message) error {

	var req common.CompanyCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return errors.Wrap(err, "error while unmarshaling company")
	}

	e.log.Info("create company event", logger.Any("event", req))

	if err := e.strg.Company().Upsert(ctx, &req); err != nil {
		e.log.Info("error while upsert company", logger.Error(err))

		return err
	}

	e.log.Info("company shop is about to create", logger.Any("event", req.Shop))

	if req.Shop == nil {
		return errors.New("error while creating shop")
	}

	if err := e.strg.Shop().Upsert(req.Shop); err != nil {
		return err
	}

	if req.Cashbox == nil {
		return errors.New("error while creating cashbox")

	}

	e.log.Info("shop cashbox is about to create", logger.Any("event", req.Cashbox))

	if err := e.strg.Cashbox().Upsert(req.Cashbox); err != nil {
		return err
	}

	return nil

}
