package repo

import (
	"context"
	"genproto/order_service"
)

type OrderI interface {
	Create(context.Context, *order_service.CreateOrderCopyRequest) error
}
