package postgres

import (
	"genproto/common"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/pkg/errors"
)

type measurementUnitRepo struct {
	db  models.DB
	log logger.Logger
}

func NewMeasurementRepo(db models.DB, log logger.Logger) repo.MeasurementI {
	return &measurementUnitRepo{
		db:  db,
		log: log,
	}
}

func (c *measurementUnitRepo) Upsert(entity *common.MeasurementUnitCopyRequest) error {
	query := `
		INSERT INTO
			"measurement_unit"
		(
			id,
			company_id,
			is_deletable,
			short_name,
			long_name,
			precision
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`

	_, err := c.db.Exec(
		query,
		entity.Id,
		entity.CompanyId,
		entity.IsDeletable,
		entity.ShortName,
		entity.LongName,
		entity.Precision,
	)
	if err != nil {
		return errors.Wrap(err, "error while upsert measurement_unit")
	}

	return nil
}

func (c *measurementUnitRepo) UpsertMany(req *common.MeasurementUnitsCopyRequest) error {

	for _, item := range req.MeasurementUnits {

		query := `
		INSERT INTO
			"measurement_unit"
		(
			id,
			company_id,
			is_deletable,
			short_name,
			long_name,
			precision
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

		_, err := c.db.Exec(
			query,
			item.Id,
			item.CompanyId,
			item.IsDeletable,
			item.ShortName,
			item.LongName,
			item.Precision,
		)
		if err != nil {
			return errors.Wrap(err, "error while create measurementunits")
		}
	}

	return nil
}

func (c *measurementUnitRepo) Delete(req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}
