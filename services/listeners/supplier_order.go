package listeners

import (
	"context"
	"fmt"
	"genproto/common"
	"genproto/inventory_service"

	"github.com/Invan2/invan_inventory_service/events/topics"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/pkg/errors"
)

func (i *inventoryService) CreateSupplierOrder(ctx context.Context, req *inventory_service.CreateSupplierOrderRequest) (*common.ResponseID, error) {
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

	res, err := tr.SupplierO().Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *inventoryService) UpsertSupplierOrderItem(ctx context.Context, req *inventory_service.CreateSupplierOrderItemRequest) (*common.ResponseID, error) {
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
	req.SoldAmount = (req.Cost * req.ExpectedAmount) * (1 - req.Discount/100)
	supplierOrderItemId, err := tr.SupplierO().UpsertOrderItem(ctx, req)
	if err != nil {
		return nil, err
	}
	res, err := tr.SupplierO().GetOrderItem(supplierOrderItemId)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.Request)
	err = i.elastic.SupplierOrderItem().UpsertOrderItem(&inventory_service.SupplierOrderItemES{
		Id:              res.Id,
		ProductId:       res.ProductId,
		Barcode:         res.Barcode,
		ExpectedAmount:  res.ExpectedAmount,
		ArrivedAmount:   res.ArrivedAmount,
		Cost:            res.Cost,
		Discount:        res.Discount,
		SoldAmount:      res.SoldAmount,
		ProductName:     res.ProductName,
		SupplierOrderId: res.SupplierOrderId,
		CompanyId:       req.Request.CompanyId,
		CreatedAt:       res.CreatedAt,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while creating order items on Elastic")
	}

	return &common.ResponseID{
		Id: res.Id,
	}, nil
}

func (i *inventoryService) GetAllSupplierOrder(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSupplierOrderResponse, error) {
	return i.strg.SupplierO().GetAll(ctx, req)
}

func (i *inventoryService) GetSupplierOrderById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierOrderByIdResponse, error) {
	return i.strg.SupplierO().GetById(ctx, req)
}

func (i *inventoryService) UpdateSupplierOrderStatus(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	return i.strg.SupplierO().UpdateSupplierOrderStatus(ctx, req, "973f4626-d0a3-4693-95a6-eb65e87b3a76")
}

func (i *inventoryService) FinishSupplierOrder(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) (*common.ResponseID, error) {

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

	err = tr.SupplierO().FinishSupplierOrder(ctx, req)
	if err != nil {
		i.log.Error(fmt.Sprintf("err occured while finishSupplierOrder: %s", err.Error()))
		return nil, err
	}

	res, err := tr.Product().GetSupplierOrderProducts(ctx, req)
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

	return &common.ResponseID{Id: req.GetSupplierOrderId()}, ctx.Err()
}

func (i *inventoryService) UpdateSupplierOrderAmount(ctx context.Context, req *inventory_service.UpdateSupplierOrderRecivedRequest) (*common.ResponseID, error) {
	return i.strg.SupplierO().UpdateSupplierOrderAmount(ctx, req)
}

func (i *inventoryService) DeleteSupplierOrder(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
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

	res, err := tr.SupplierO().Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *inventoryService) GetAllSupplierOrderItems(ctx context.Context, entity *inventory_service.GetAllSupplierOrderItemsRequest) (res *inventory_service.GetAllSupplierOrderItemsResponse, err error) {
	return i.elastic.SupplierOrderItem().GetAllOrderItems(entity)

}

func (i *inventoryService) DeleteSupplierOrderItemById(ctx context.Context, req *common.RequestID) (*common.Empty, error) {

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

	err = tr.SupplierO().DeleteItemById(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = i.elastic.SupplierOrderItem().DeleteSupplierOrderItemById(req)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}
