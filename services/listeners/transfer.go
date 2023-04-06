package listeners

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
)

func (i *inventoryService) CreateTransfer(ctx context.Context, in *inventory_service.CreateTransferRequest) (*common.ResponseID, error) {
	return nil, nil
}
func (i *inventoryService) GetAllTransfer(ctx context.Context, in *inventory_service.GetAllTransferRequest) (*inventory_service.GetAllTransferResponse, error) {
	return nil, nil
}
func (i *inventoryService) GetAllTransferItems(ctx context.Context, in *inventory_service.GetAllTransferItemsRequest) (*inventory_service.GetAllTransferItemsResponse, error) {
	return nil, nil
}
func (i *inventoryService) AddItemToTransfer(ctx context.Context, in *inventory_service.AddItemToTransferRequest) (*common.ResponseID, error) {
	return nil, nil
}
