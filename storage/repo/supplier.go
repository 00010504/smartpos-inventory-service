package repo

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
)

type SupplierI interface {
	Create(ctx context.Context, req *inventory_service.CreateSupplierRequest) (*common.ResponseID, error)
	GetAll(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSuppliersResponse, error)
	GetById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierByIdResponse, error)
	Update(ctx context.Context, req *inventory_service.UpdateSupplierRequest) (*common.ResponseID, error)
	Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
}
