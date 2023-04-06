package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/pkg/errors"
)

type cashboxRepo struct {
	db  models.DB
	log logger.Logger
}

func NewCashboxRepo(db models.DB, log logger.Logger) repo.CashboxI {
	return &cashboxRepo{
		db:  db,
		log: log,
	}
}

func (c *cashboxRepo) Upsert(entity *common.CashboxCreatedModel) error {

	query := `
		INSERT INTO
			"cashbox"
		(
			id,
			shop_id,
			company_id,
			title,
			created_by
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`

	_, err := c.db.Exec(
		query,
		entity.Id,
		entity.ShopId,
		entity.Request.CompanyId,
		entity.Title,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert cashbox")
	}

	return nil
}

func (c *cashboxRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}
