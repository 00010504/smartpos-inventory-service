package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"genproto/common"
	"genproto/inventory_service"
	pbu "genproto/inventory_service"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/helper"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type importRepo struct {
	db  models.DB
	log logger.Logger
}

func NewImportRepo(db models.DB, log logger.Logger) repo.ImportI {
	return &importRepo{
		db:  db,
		log: log,
	}
}

func (i *importRepo) Create(ctx context.Context, req *pbu.CreateImportRequest) error {
	err := i.db.QueryRow(
		`insert into import(
			id,
			external_id,
			name,
			shop_id,
			quantity,
			total_price,
			supplier,
			type,
			created_by,
			company_id,
			completed_by,
			status_id
		)
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,	$11, $12)
		returning id`,
		req.GetId(),
		strconv.Itoa(int(req.GetExternalId())),
		req.GetName(),
		req.GetShopId(),
		req.GetQuantity(),
		req.GetTotalPrice(),
		req.GetSupplier(),
		req.GetType(),
		req.GetCreatedBy(),
		req.GetCompanyId(),
		req.GetCompletedBy(),
		"0b123839-be70-4372-866f-fce1c59f6ed0").Scan(&req.Id)
	if err != nil {
		i.log.Error(fmt.Sprintf("err occured while create import: %s", err.Error()))
		return err
	}
	return nil
}

func (ii *importRepo) CreateImportItems(ctx context.Context, req *pbu.CreateImportRequest) error {

	var (
		productValues                = make([]interface{}, 0)
		inmportItemsValues           = make([]interface{}, 0)
		productInsertSubQuery string = ""
		importTiemSubQuery    string = ""
	)

	productInsertQuery := `
	INSERT INTO
		"product"
	(
		id,
		name,
		is_marking,
		sku,
		company_id,
		created_by
	)
	VALUES 
	`

	conflict := `
	ON CONFLICT (sku, company_id, deleted_at) DO NOTHING
	`

	importTiemQuery := `insert into import_item(
		id,
		name,
		import_id,
		supply_price,
		sale_price,
		barcode,
		quantity,
		sku,
		measurement_unit_id,
		product_id
	)
	VALUES`
	//create product many
	for i, item := range req.GetItems() {
		productInsertSubQuery += "(?, ?, ?, ?, ?, ?),"
		productValues = append(productValues,
			uuid.New().String(),
			item.Name,
			false,
			item.GetSku(),
			req.GetCompanyId(),
			req.GetCreatedBy(),
		)

		importTiemSubQuery += "(?, ?, ?, ?, ?, ?, ?, ?, ?,  (SELECT id FROM product as p WHERE p.sku = ? and p.company_id = ? and p.deleted_at = 0)),"
		inmportItemsValues = append(inmportItemsValues,
			item.GetId(),
			item.GetName(),
			req.GetId(),
			item.GetSupplyPrice(),
			item.GetSalePrice(),
			item.GetBarcode(),
			item.GetQuantity(),
			item.GetSku(),
			item.GetMeasuramentUnit().GetId(),
			item.GetSku(),
			req.GetCompanyId(),
		)

		if i%1000 == 999 && i != 0 || (i+1) == len(req.GetItems()) {
			fmt.Println("Product Insert query: ", strings.Count(productInsertSubQuery, "(?, ?, ?, ?, ?, ?)"), " length : ", len(productValues))
			productInsertSubQuery = strings.TrimSuffix(productInsertSubQuery, ",")
			productInsertSubQuery = helper.ReplaceSQL(productInsertSubQuery, "?")

			stmt, err := ii.db.Prepare(productInsertQuery + productInsertSubQuery + conflict)
			if err != nil {
				return errors.Wrap(err, "error while insertMany products. Prepare")
			}
			defer stmt.Close()

			_, err = stmt.Exec(productValues...)
			if err != nil {
				return errors.Wrap(err, "error while insertMany products. Exec")
			}

			importTiemSubQuery = strings.TrimSuffix(importTiemSubQuery, ",")
			importTiemSubQuery = helper.ReplaceSQL(importTiemSubQuery, "?")

			stmt1, err := ii.db.Prepare(importTiemQuery + importTiemSubQuery)
			if err != nil {
				return errors.Wrap(err, "error while insertMany import items. Prepare")
			}
			defer stmt1.Close()

			_, err = stmt1.Exec(inmportItemsValues...)
			if err != nil {
				return errors.Wrap(err, "error while insertMany import item. Exec")
			}

			productValues = nil
			inmportItemsValues = nil
			importTiemSubQuery = ""
			productInsertSubQuery = ""
		}
	}
	return nil
}

func (i *importRepo) GetAll(ctx context.Context, req *common.SearchRequest) (res *pbu.GetAllImportRes, err error) {
	var count int32
	res = &pbu.GetAllImportRes{}

	countQuery := `SELECT count(1) FROM import where company_id = $1`
	rows, err := i.db.Query(countQuery, req.GetRequest().GetCompanyId())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT
		i.id,
		i.name,
		i.quantity,
		i.status_id,
		i.total_price,
		i.type,
		i.supplier,
		i.external_id,
		sh.id,
		sh.name,
		c.id, 
		c.name, 
		u.id,
		COALESCE(u.first_name,''),
		COALESCE(u.last_name,''),
		TO_CHAR(
			i.created_at,
			'YYYY-MM-DD HH:MI:SS'
		) created_atC
	FROM
		"import" AS i 
	LEFT JOIN "shop" sh ON sh.id = i.shop_id 
	LEFT JOIN "user" u ON u.id = i.created_by 
	LEFT JOIN "user" cu ON cu.id = i.completed_by 
	LEFT JOIN "company" c ON c.id = i.company_id 
	WHERE i.company_id = $1 AND i.deleted_at = 0`
	rows, err = i.db.Query(query, req.GetRequest().GetCompanyId())
	if err != nil {
		i.log.Error(fmt.Sprintf("err occured while get rows for get all import: %s", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			imp         pbu.GetAllImportResModel
			shop        common.ShortShop
			company     common.ShortCompany
			createdUser common.ShortUser
			// completedUser common.ShortUser
			external_id string
		)
		err = rows.Scan(
			&imp.Id,
			&imp.Name,
			&imp.Quantity,
			&imp.Status,
			&imp.TotalPrice,
			&imp.Type,
			&imp.Supplier,
			&external_id,
			&shop.Id,
			&shop.Name,
			&company.Id,
			&company.Name,
			&createdUser.Id,
			&createdUser.FirstName,
			&createdUser.LastName,
			&imp.CreatedAt,
		)
		if err != nil {
			i.log.Error(fmt.Sprintf("err occured while scan for get all import: %s", err.Error()))
			return nil, err
		}

		exid, err := strconv.Atoi(external_id)
		if err != nil {
			exid = 0
		}
		imp.Company = &company
		imp.Shop = &shop
		imp.CreatedBy = &createdUser
		imp.CompletedBy = &createdUser
		imp.ExternalId = int32(exid)
		res.Imports = append(res.Imports, &imp)
	}
	res.Count = count
	return
}

func (i *importRepo) GetImportByID(ctx context.Context, req *common.RequestID) (*inventory_service.ImportById, error) {

	var res inventory_service.ImportById
	query := `
		SELECT
			i.id,
			i.external_id,
			i.status_id,
			i.name,
			i.total_price,
			i.shop_id,
			i.company_id,
			i.created_by,
			ii.supply_price,
			ii.sale_price
		FROM "import" i
		LEFT JOIN "import_item" ii ON ii.import_id = i.id
		WHERE
		i.id = $1 and i.deleted_at = 0
	`

	err := i.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.ExternalId,
		&res.Status,
		&res.Name, &res.TotalPrice, &res.ShopId, &res.CompanyId, &res.CreatedBy, &res.SupplyPrice, &res.RetailPrice)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (i *importRepo) FinishImport(ctx context.Context, req *inventory_service.FinishImportReq) error {
	query := `update import set
		status_id = $2
		WHERE id = $1 `
	_, err := i.db.Exec(
		query,
		req.ImportId,
		"0b123839-be70-4372-866f-fce1c59f6ed1",
	)
	if err != nil {
		return errors.Wrap(err, "error while finish import")
	}
	return nil
}

func (i *importRepo) Upsert(ctx context.Context, req *common.CreateProductCopyRequest) error {
	var count int = 0
	query := `
	SELECT
		count(*)
	FROM "product" p
	WHERE p.id = $1 AND company_id = $2 AND p.deleted_at = 0 AND p.updated_at = 0
	`

	err := i.db.QueryRow(query, req.Id, req.Request.CompanyId).Scan(
		&count)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return err
		}
	}
	external_id := helper.RandStringRunes(6, "number")
	if count == 1 {
		barcode := ""
		if len(req.GetBarcode()) > 0 {
			barcode = req.Barcode[0]
		}
		for _, v := range req.ShopMeasurementValues {
			var importId = uuid.New().String()
			var productid string
			// CREATE IMPORT_ITEM
			err = i.db.QueryRow(
				`insert into import_item(
					id,
					name,
					import_id,
					supply_price,
					sale_price,
					barcode,
					quantity,
					sku,
					measurement_unit_id,
					product_id
				)
				values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
				returning product_id`,
				uuid.New().String(),
				req.GetName(),
				importId,
				v.GetSupplyPrice(),
				v.GetRetailPrice(),
				barcode,
				v.GetInStock(),
				req.GetSku(),
				req.GetMeasurementUnitId(),
				req.Id).Scan(&productid)
			if err != nil {
				i.log.Error(fmt.Sprintf("err occured while create import_item on upsert product: %s", err.Error()))
				return err
			}

			// CREATE IMPORT
			err := i.db.QueryRow(
				`insert into import(
					id,
					external_id,
					name,
					shop_id,
					quantity,
					total_price,
					type,
					created_by,
					company_id,
					completed_by,
					status_id
				)
				values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,	$11)
				returning id`,
				importId,
				external_id,
				"New import - "+req.GetName(),
				v.GetShopId(),
				v.GetInStock(),
				v.GetInStock()*v.GetSupplyPrice(),
				req.GetProductTypeId(),
				req.GetRequest().GetUserId(),
				req.GetRequest().GetCompanyId(),
				req.GetRequest().GetUserId(),
				"0b123839-be70-4372-866f-fce1c59f6ed1").Scan(&importId)
			if err != nil {
				i.log.Error(fmt.Sprintf("err occured while create import on upsert product: %s", err.Error()))
				return err
			}
		}
	}
	return nil
}
