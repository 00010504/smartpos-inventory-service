package listeners

import (
	"context"
	"fmt"
	"genproto/common"
	"genproto/inventory_service"

	"github.com/Invan2/invan_inventory_service/events/topics"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
)

func (i *inventoryService) CreateImport(ctx context.Context, req *inventory_service.CreateImportRequest) (*common.ResponseID, error) {
	i.log.Info("CreateImport request", logger.Any("id: ", req.Id))

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

	err = tr.Import().Create(ctx, req)
	if err != nil {
		i.log.Error(fmt.Sprintf("error occured while creating imports : %s", err.Error()))
		return nil, err
	}

	err = tr.Import().CreateImportItems(ctx, req)
	if err != nil {
		i.log.Error(fmt.Sprintf("error occured while creating imports items : %s", err.Error()))
		return nil, err
	}
	return &common.ResponseID{Id: req.GetId()}, ctx.Err()
}

func (i *inventoryService) FinishImport(ctx context.Context, req *inventory_service.FinishImportReq) (*common.ResponseID, error) {
	i.log.Info("FinishImport request", logger.Any("req: ", req))

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

	err = tr.Import().FinishImport(ctx, req)
	if err != nil {
		i.log.Error(fmt.Sprintf("err occured while finishImport: %s", err.Error()))
		return nil, err
	}

	res, err := tr.Product().GetImportProducts(ctx, req)
	if err != nil {
		i.log.Error(fmt.Sprintf("err occured while get GetImportProducts: %s", err.Error()))
		return nil, err
	}

	var from int = 0
	for j := range res {
		if j > 0 && (j%1000 == 0 || j == len(res)-1) {
			if j == len(res)-1 {
				err = i.kafka.Push(topics.CatalogCreateMultipleProducts, &common.CreateImportProductsModel{Products: res[from:]})
			} else {
				err = i.kafka.Push(topics.CatalogCreateMultipleProducts, &common.CreateImportProductsModel{Products: res[from:j]})
			}
			if err != nil {
				i.log.Error(fmt.Sprintf("error occured while pushing event of CatalogCreateMultipleProducts : %s", err.Error()))
				return nil, err
			}
			i.log.Info("FinishImport message sent", logger.Any("message: ", len(res[from:j])))
			from = j
		}
		if len(res) == 1 {
			err = i.kafka.Push(topics.CatalogCreateMultipleProducts, &common.CreateImportProductsModel{Products: res})
			if err != nil {
				i.log.Error(fmt.Sprintf("error occured while pushing event of CatalogCreateMultipleProducts : %s", err.Error()))
				return nil, err
			}
			i.log.Info("FinishImport message sent", logger.Any("message: ", len(res)))
			from = j
		}
	}
	return &common.ResponseID{Id: req.GetImportId()}, ctx.Err()
}

func (i *inventoryService) GetAllImport(ctx context.Context, req *common.SearchRequest) (res *inventory_service.GetAllImportRes, err error) {
	i.log.Info("getting all import", logger.Any("req: ", req))
	res, err = i.strg.Import().GetAll(ctx, req)
	if err != nil {
		i.log.Error("error getting all import", logger.Any("err: ", err.Error()))
		return nil, err
	}
	return res, nil
}

func (i *inventoryService) GetImportByID(ctx context.Context, req *common.RequestID) (*inventory_service.ImportById, error) {
	i.log.Info("getting all GetImportByID by id", logger.Any("req: ", req))
	res, err := i.strg.Import().GetImportByID(ctx, req)
	if err != nil {
		i.log.Error("error while getting all GetImportByID by id", logger.Any("error: ", err.Error()))
		return nil, err
	}
	return res, nil

}

func (i *inventoryService) GetProductHistory(ctx context.Context, req *inventory_service.GetProductHistoryReq) (*inventory_service.GetProductHistoryRes, error) {
	i.log.Info("getting all GetProductHistory by id", logger.Any("req: ", req))
	res, err := i.strg.Product().GetProductHistoryById(ctx, req)
	if err != nil {
		i.log.Error("error while getting all GetProductHistory by id", logger.Any("error: ", err.Error()))
		return nil, err
	}
	return res, nil
}
