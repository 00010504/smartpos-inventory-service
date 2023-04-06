package listeners

import (
	"context"
	"genproto/common"
	"genproto/order_service"

	"github.com/Invan2/invan_inventory_service/events/topics"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
)

func (i *inventoryService) CreateOrder(ctx context.Context, req *order_service.CreateOrderCopyRequest) (*common.ResponseID, error) {

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

	err = tr.Order().Create(ctx, req)
	if err != nil {
		i.log.Error("error occured while CreateOrder", logger.Any("tr.Order().Create", err.Error()))
		return nil, err
	}

	items, err := tr.Product().GetOrderedProducts(ctx, req)
	if err != nil {
		i.log.Error("error occured while CreateOrder", logger.Any("tr.Product().GetOrderedProducts", err.Error()))
		return nil, err
	}

	shop, err := tr.Shop().GetById(&common.RequestID{Id: req.ShopId})
	if err != nil {
		i.log.Error("error occured while CreateOrder", logger.Any("tr.Shop().GetById", err.Error()))
		return nil, err
	}

	err = i.kafka.Push(topics.OrdertCreateTopic, &common.OrderCopyRequest{
		OrderId:   req.OrderId,
		CompanyId: req.CompanyId,
		ShopId:    req.ShopId,
		ShopName:  shop.Title,
		Request:   req.Request,
		Items:     items,
	})
	if err != nil {
		i.log.Error("error occured while CreateOrder", logger.Any("i.kafka.Push", err.Error()))
		return nil, err
	}

	return &common.ResponseID{Id: req.OrderId}, nil
}
