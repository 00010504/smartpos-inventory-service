package repo

import (
	"genproto/common"
	"genproto/inventory_service"
)

type SupplierOrderItemESI interface {
	UpsertOrderItem(req *inventory_service.SupplierOrderItemES) error
	CreateMultiple(req *inventory_service.CreateSupplierOrderElasticRequest) error
	GetAllOrderItems(entity *inventory_service.GetAllSupplierOrderItemsRequest) (*inventory_service.GetAllSupplierOrderItemsResponse, error)
	DeleteSupplierOrderItemById(req *common.RequestID) (*common.Empty, error)
}
