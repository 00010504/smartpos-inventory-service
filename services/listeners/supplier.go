package listeners

import (
	"context"
	"genproto/common"
	"genproto/inventory_service"
)

func (i *inventoryService) CreateSupplier(ctx context.Context, req *inventory_service.CreateSupplierRequest) (*common.ResponseID, error) {

	tr, err := i.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := tr.Supplier().Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *inventoryService) GetAllSupplier(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSuppliersResponse, error) {

	return i.strg.Supplier().GetAll(ctx, req)

}

func (i *inventoryService) GetSupplierById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierByIdResponse, error) {
	return i.strg.Supplier().GetById(ctx, req)
}

func (i *inventoryService) UpdateSupplier(ctx context.Context, req *inventory_service.UpdateSupplierRequest) (*common.ResponseID, error) {
	tr, err := i.strg.WithTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tr.Rollback()
		} else {
			_ = tr.Commit()
		}
	}()

	res, err := tr.Supplier().Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *inventoryService) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	return i.strg.Supplier().Delete(ctx, req)
}
