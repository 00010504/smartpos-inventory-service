package repo

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
	"genproto/order_service"
)

type ProductI interface {
	Upsert(ctx context.Context, entity *common.CreateProductCopyRequest) error
	Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
	GetImportProducts(ctx context.Context, req *inventory_service.FinishImportReq) ([]*common.CreateProductCopyRequest, error)
	GetSupplierOrderProducts(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) (res []*common.CreateProductCopyRequest, err error)
	GetProductHistoryById(ctx context.Context, req *inventory_service.GetProductHistoryReq) (res *inventory_service.GetProductHistoryRes, err error)
	GetOrderedProducts(ctx context.Context, req *order_service.CreateOrderCopyRequest) (res []*common.CreateOrderItemCopyRequest, err error)
}
