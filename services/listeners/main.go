package listeners

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
	"genproto/order_service"

	"github.com/Invan2/invan_inventory_service/events"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage"
)

type inventoryService struct {
	strg    storage.StoragePgI
	log     logger.Logger
	kafka   events.PubSubServer
	elastic storage.StorageES
}

type InventoryService interface {
	CreateImport(context.Context, *inventory_service.CreateImportRequest) (*common.ResponseID, error)
	FinishImport(context.Context, *inventory_service.FinishImportReq) (*common.ResponseID, error)
	CreateOrder(ctx context.Context, req *order_service.CreateOrderCopyRequest) (*common.ResponseID, error)
	GetImportByID(ctx context.Context, req *common.RequestID) (*inventory_service.ImportById, error)
	GetAllImport(context.Context, *common.SearchRequest) (*inventory_service.GetAllImportRes, error)

	// supplier
	CreateSupplier(ctx context.Context, req *inventory_service.CreateSupplierRequest) (*common.ResponseID, error)
	GetAllSupplier(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSuppliersResponse, error)
	GetSupplierById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierByIdResponse, error)
	UpdateSupplier(ctx context.Context, req *inventory_service.UpdateSupplierRequest) (*common.ResponseID, error)
	Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)

	// supplierOrder
	CreateSupplierOrder(ctx context.Context, req *inventory_service.CreateSupplierOrderRequest) (*common.ResponseID, error)
	UpsertSupplierOrderItem(ctx context.Context, req *inventory_service.CreateSupplierOrderItemRequest) (*common.ResponseID, error)
	GetAllSupplierOrder(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSupplierOrderResponse, error)
	GetAllSupplierOrderItems(ctx context.Context, entity *inventory_service.GetAllSupplierOrderItemsRequest) (res *inventory_service.GetAllSupplierOrderItemsResponse, err error)
	GetSupplierOrderById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierOrderByIdResponse, error)
	UpdateSupplierOrderStatus(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
	FinishSupplierOrder(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) (*common.ResponseID, error)
	UpdateSupplierOrderAmount(ctx context.Context, req *inventory_service.UpdateSupplierOrderRecivedRequest) (*common.ResponseID, error)
	DeleteSupplierOrder(ctx context.Context, req *common.RequestID) (*common.ResponseID, error)
	DeleteSupplierOrderItemById(ctx context.Context, req *common.RequestID) (*common.Empty, error)

	//Product
	GetProductHistory(ctx context.Context, req *inventory_service.GetProductHistoryReq) (*inventory_service.GetProductHistoryRes, error)

	CreateTransfer(ctx context.Context, in *inventory_service.CreateTransferRequest) (*common.ResponseID, error)
	GetAllTransfer(ctx context.Context, in *inventory_service.GetAllTransferRequest) (*inventory_service.GetAllTransferResponse, error)
	GetAllTransferItems(ctx context.Context, in *inventory_service.GetAllTransferItemsRequest) (*inventory_service.GetAllTransferItemsResponse, error)
	AddItemToTransfer(ctx context.Context, in *inventory_service.AddItemToTransferRequest) (*common.ResponseID, error)
}

func NewInventoryService(strg storage.StoragePgI, log logger.Logger, kafka events.PubSubServer, elastic storage.StorageES) InventoryService {
	return &inventoryService{
		strg:    strg,
		log:     log,
		kafka:   kafka,
		elastic: elastic,
	}
}
