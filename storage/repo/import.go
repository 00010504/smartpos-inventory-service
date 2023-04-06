package repo

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
	pbu "genproto/inventory_service"
)

type ImportI interface {
	Upsert(ctx context.Context, entity *common.CreateProductCopyRequest) error
	Create(ctx context.Context, req *pbu.CreateImportRequest) error
	GetAll(ctx context.Context, req *common.SearchRequest) (*pbu.GetAllImportRes, error)
	CreateImportItems(ctx context.Context, req *pbu.CreateImportRequest) error
	GetImportByID(ctx context.Context, req *common.RequestID) (*inventory_service.ImportById, error)
	FinishImport(ctx context.Context, req *inventory_service.FinishImportReq) error
}
