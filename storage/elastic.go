package storage

import (
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/elastic"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/elastic/go-elasticsearch/v8"
)

type storageES struct {
	db                    *elasticsearch.Client
	log                   logger.Logger
	supplierOrderItemRepo repo.SupplierOrderItemESI
}

type StorageES interface {
	SupplierOrderItem() repo.SupplierOrderItemESI
}

func NewStorageES(log logger.Logger, db *elasticsearch.Client) StorageES {
	return &storageES{
		db:                    db,
		log:                   log,
		supplierOrderItemRepo: elastic.NewSupplierOrderRepo(log, db),
	}
}

func (s *storageES) SupplierOrderItem() repo.SupplierOrderItemESI {
	return s.supplierOrderItemRepo
}
