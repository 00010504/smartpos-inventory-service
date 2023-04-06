package handlers

import (
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage"
)

type EventHandler struct {
	log  logger.Logger
	strg storage.StoragePgI
}

func NewHandler(log logger.Logger, strg storage.StoragePgI) *EventHandler {
	return &EventHandler{
		log:  log,
		strg: strg,
	}
}
