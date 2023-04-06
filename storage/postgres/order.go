package postgres

import (
	"context"
	"genproto/order_service"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/pkg/errors"
)

type orderRepo struct {
	db  models.DB
	log logger.Logger
}

func NewOrderRepo(db models.DB, log logger.Logger) repo.OrderI {
	return &orderRepo{
		db:  db,
		log: log,
	}
}

func (i *orderRepo) Create(ctx context.Context, req *order_service.CreateOrderCopyRequest) error {

	query := `
		INSERT INTO
			"order"
		(
			id,
			external_id,
			status,
			total_price,
			total_discount_price,
			product_sort_count,
			shop_id,
			cashbox_id,
			company_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := i.db.Exec(
		query,
		req.OrderId,
		req.ExternalId,
		req.Status,
		req.TotalPrice,
		req.TotalDiscountPrice,
		req.ProductSortCount,
		req.ShopId,
		req.CashboxId,
		req.CompanyId,
		req.CreatedBy,
	)
	if err != nil {
		return errors.Wrap(err, "error while create order")
	}

	// create order items
	for _, item := range req.Items {

		query = `
			INSERT INTO
				"order_item"
			(
				id,
				price,
				value,
				total_price,
				order_id,
				product_id,
				created_by
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`

		_, err := i.db.Exec(
			query,
			item.Id,
			item.Price,
			item.Value,
			item.TotalPrice,
			req.OrderId,
			item.ProductId,
			item.CreatedBy,
		)
		if err != nil {
			return errors.Wrap(err, "error while create order items")
		}
	}

	return nil
}
