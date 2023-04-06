package postgres

import (
	"genproto/common"
	"genproto/corporate_service"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/pkg/errors"
)

type shopRepo struct {
	db  models.DB
	log logger.Logger
}

func NewShopRepo(db models.DB, log logger.Logger) repo.ShopI {
	return &shopRepo{db: db, log: log}
}

func (s *shopRepo) Upsert(entity *common.ShopCreatedModel) error {

	query := `
	INSERT INTO
		"shop"
	(
		id,
		name,
		company_id,
		created_by
	)
	VALUES (
		$1,
		$2,
		$3,
		$4
	) ON CONFLICT (id) DO
	UPDATE
		SET
		"name" = $2,
		"company_id" = $3
		;
`

	_, err := s.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.Request.CompanyId,
		entity.Request.UserId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert shop")
	}

	return nil
}

func (p *shopRepo) GetById(entity *common.RequestID) (*corporate_service.Shop, error) {

	var (
		shop corporate_service.Shop
	)

	var query = `
		SELECT 
			s.id,
			s.name
		FROM "shop" s
		WHERE 
			s.id = $1 and s.deleted_at = 0
	`
	err := p.db.QueryRow(query, entity.Id).Scan(
		&shop.Id,
		&shop.Title,
	)

	if err != nil {
		return nil, errors.Wrap(err, "error while getting shop")
	}

	return &shop, nil
}
