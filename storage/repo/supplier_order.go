package repo

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
)

type SupplierOrderI interface {
	Create(ctx context.Context, req *inventory_service.CreateSupplierOrderRequest) (*common.ResponseID, error)
	UpsertOrderItem(ctx context.Context, req *inventory_service.CreateSupplierOrderItemRequest) (string, error)
	GetOrderItem(supplierOrderItemId string) (*inventory_service.CreateSupplierOrderItemElasticResponse, error)
	GetAll(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSupplierOrderResponse, error)
	GetById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierOrderByIdResponse, error)
	UpdateSupplierOrderStatus(ctx context.Context, req *common.RequestID, supplierOrderStatus string) (*common.ResponseID, error)
	FinishSupplierOrder(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) error
	UpdateSupplierOrderAmount(ctx context.Context, req *inventory_service.UpdateSupplierOrderRecivedRequest) (*common.ResponseID, error)
	Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
	DeleteItemById(ctx context.Context, req *common.RequestID) error
}
