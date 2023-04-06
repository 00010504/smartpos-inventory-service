package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"genproto/common"
	"genproto/inventory_service"
	"genproto/order_service"
	"strings"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/helper"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/pkg/errors"
)

type productRepo struct {
	db  models.DB
	log logger.Logger
}

func NewProductRepo(log logger.Logger, db models.DB) repo.ProductI {
	return &productRepo{
		db:  db,
		log: log,
	}
}

func (p *productRepo) Upsert(ctx context.Context, entity *common.CreateProductCopyRequest) error {

	var values = []interface{}{}
	query := `
		INSERT INTO
			"product"
		(
			id,
			name,
			is_marking,
			sku,
			mxik_code,
			company_id
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO
		UPDATE
		SET 
		name = $2,
		is_marking = $3,
		sku = $4,
		mxik_code = $5,
		updated_at = extract(epoch from now())::bigint;
	`

	res, err := p.db.Exec(
		query,
		entity.Id,
		entity.Name,
		entity.IsMarking,
		entity.Sku,
		entity.MxikCode,
		entity.Request.CompanyId,
	)
	if err != nil {
		return errors.Wrap(err, "error while insert product")
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}
	// insert measurement values
	if len(entity.ShopMeasurementValues) > 0 {

		values = []interface{}{}

		query = `
			INSERT INTO
				"measurement_values"
			(
				shop_id,
				product_id,
				is_aviable,
				in_stock
			)
			VALUES 
		`

		for _, value := range entity.ShopMeasurementValues {
			query += "(?, ?, ?, ?),"
			values = append(values,
				value.ShopId,
				entity.Id,
				value.IsAviable,
				value.InStock,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += `
		ON CONFLICT (product_id, shop_id) DO
		UPDATE
			SET
			is_aviable = excluded.is_aviable,
			in_stock = excluded.in_stock;
		`
		stmt, err := p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product measurement_values")
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			stmt.Close()
			return errors.Wrap(err, "error while insert product measurement_values")
		}

		stmt.Close()

		values = []interface{}{}
		query = `
			INSERT INTO
				"shop_price"
			(
				supply_price,
				min_price,
				max_price,
				retail_price,
				whole_sale_price,
				shop_id,
				product_id
			)
			VALUES 
		`

		for _, value := range entity.ShopMeasurementValues {
			query += "(?, ?, ?, ?, ?, ?, ?),"
			values = append(values,
				value.SupplyPrice,
				value.MinPrice,
				value.MaxPrice,
				value.RetailPrice,
				value.WholeSalePrice,
				value.ShopId,
				entity.Id,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += `
		ON CONFLICT (product_id, shop_id) DO
		UPDATE
			SET
			supply_price = excluded.supply_price,
			min_price = excluded.min_price,
			max_price = excluded.max_price,
			retail_price = excluded.retail_price,
			whole_sale_price = excluded.whole_sale_price;
		`

		stmt, err = p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product shop_price. Prepare")
		}
		_, err = stmt.Exec(values...)
		if err != nil {
			return errors.Wrap(err, "error while insert product shop_price. Exec")
		}

		stmt.Close()
	}

	// insert barcodes
	if len(entity.Barcode) > 0 {
		values = []interface{}{}
		query = `
			INSERT INTO
				"product_barcode"
			(
				barcode,
				product_id
			)
			VALUES 
		`

		for _, barcode := range entity.Barcode {
			query += "(?, ?),"
			values = append(values,
				barcode,
				entity.Id,
			)
		}

		query = strings.TrimSuffix(query, ",")
		query = helper.ReplaceSQL(query, "?")

		query += `
		ON CONFLICT (barcode, product_id) DO NOTHING
		`

		stmt, err := p.db.Prepare(query)
		if err != nil {
			return errors.Wrap(err, "error while insert product barcodes. Prepare")
		}

		_, err = stmt.Exec(values...)
		if err != nil {
			stmt.Close()
			return errors.Wrap(err, "error while insert product barcodes. Exec")
		}

		stmt.Close()
	}

	return nil
}

func (p *productRepo) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	return nil, nil
}

func (p *productRepo) GetImportProducts(ctx context.Context, req *inventory_service.FinishImportReq) (res []*common.CreateProductCopyRequest, err error) {

	getProductByImportId := `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.is_marking,
		ii.measurement_unit_id,
		ii.barcode,
		ii.supply_price,
		ii.sale_price,
		sum(remainder_count) as quantity,
		i.shop_id,
		p.type
	FROM
		"product" as p
	JOIN "import_item" AS ii ON ii.product_id = p.id
	JOIN "import" AS i ON i.id = ii.import_id
	JOIN "product_queue" AS pq ON p.id = pq.product_id AND i.shop_id = pq.shop_id
	WHERE
		i.id = $1 and i.company_id = $2
	GROUP BY 
		p.id,
		p.sku,
		p.name,
		p.is_marking,
		ii.measurement_unit_id,
		ii.barcode,
		ii.supply_price,
		ii.sale_price,
		i.shop_id,
		p.type`
	// err = p.db.QueryRow(ctx, getProductByImportId, req.GetImportId(), req.GetRequest().GetCompanyId()).Scan(&res)
	rows, err := p.db.Query(getProductByImportId, req.GetImportId(), req.GetRequest().GetCompanyId())
	if err != nil {
		p.log.Error(fmt.Sprintf("err occured while get rows for GetImportProducts: %s", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			product                    common.CreateProductCopyRequest
			barcode                    string
			commonShopMeasurementValue common.CommonShopMeasurementValue
		)
		err = rows.Scan(
			&product.Id,
			&product.Sku,
			&product.Name,
			&product.IsMarking,
			&product.MeasurementUnitId,
			&barcode,
			&commonShopMeasurementValue.SupplyPrice,
			&commonShopMeasurementValue.RetailPrice,
			&commonShopMeasurementValue.InStock,
			&commonShopMeasurementValue.ShopId,
			&product.ProductTypeId,
		)
		if err != nil {
			p.log.Error(fmt.Sprintf("err occured while scan rows for GetImportProducts: %s", err.Error()))
			return nil, err
		}

		// product.MeasurementUnitId = "b03e13f8-77c7-4b73-b282-3a2c50d18ca6"
		product.ShopMeasurementValues = append(product.ShopMeasurementValues, &commonShopMeasurementValue)
		product.Barcode = append(product.Barcode, barcode)
		product.Request = req.Request

		res = append(res, &product)
	}
	return
}

func (p *productRepo) GetSupplierOrderProducts(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) (res []*common.CreateProductCopyRequest, err error) {

	getProductByImportId := `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.is_marking,
		p.measurement_unit_id,
		pb.barcode,
		sum(remainder_count) as quantity,
		so.shop_id,
		p.type
	FROM
		"product" as p
	JOIN "supplier_order_item" AS soi ON soi.product_id = p.id
	JOIN "supplier_order" AS so ON so.id = soi.supplier_order_id
	JOIN "product_barcode" AS pb on pb.product_id = p.id
	JOIN "product_queue" AS pq ON p.id = pq.product_id AND so.shop_id = pq.shop_id
	WHERE
		so.id = $1 and so.company_id = $2
	GROUP BY 
		p.id,
		p.sku,
		p.name,
		p.is_marking,
		p.measurement_unit_id,
		pb.barcode,
		so.shop_id,
		p.type`
	// err = p.db.QueryRow(ctx, getProductByImportId, req.GetImportId(), req.GetRequest().GetCompanyId()).Scan(&res)
	rows, err := p.db.Query(getProductByImportId, req.GetSupplierOrderId(), req.GetRequest().GetCompanyId())
	if err != nil {
		p.log.Error(fmt.Sprintf("err occured while get rows for GetSupplierOrderProducts: %s", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			product                    common.CreateProductCopyRequest
			barcode                    string
			commonShopMeasurementValue common.CommonShopMeasurementValue
		)
		err = rows.Scan(
			&product.Id,
			&product.Sku,
			&product.Name,
			&product.IsMarking,
			&product.MeasurementUnitId,
			&barcode,
			&commonShopMeasurementValue.InStock,
			&commonShopMeasurementValue.ShopId,
			&product.ProductTypeId,
		)
		if err != nil {
			p.log.Error(fmt.Sprintf("err occured while scan rows for GetSupplierOrderProducts: %s", err.Error()))
			return nil, err
		}

		product.ShopMeasurementValues = append(product.ShopMeasurementValues, &commonShopMeasurementValue)
		product.Barcode = append(product.Barcode, barcode)
		product.Request = req.Request

		res = append(res, &product)
	}
	return
}

func (p *productRepo) GetProductHistoryById(ctx context.Context, req *inventory_service.GetProductHistoryReq) (res *inventory_service.GetProductHistoryRes, err error) {
	var count int32
	res = &inventory_service.GetProductHistoryRes{}

	getProductHistoryById := `
	SELECT
		p.product_id,
		p.shop_id,
		p.quantity,
		p.source_id,
		p.created_at,
		p.external_id,
		sh.name
	FROM
		"product_history" as p
	JOIN "shop" AS sh ON sh.id = p.shop_id
	WHERE
		p.product_id = $1`
	rows, err := p.db.Query(getProductHistoryById, req.GetProductId())
	if err != nil {
		p.log.Error(fmt.Sprintf("err occured while get rows for GetProductHistoryById: %s", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			productHistory inventory_service.GetProductHistoryResModel
		)
		err = rows.Scan(
			&productHistory.ProductId,
			&productHistory.ShopId,
			&productHistory.Quantity,
			&productHistory.SourceId,
			&productHistory.CreatedAt,
			&productHistory.ExternalId,
			&productHistory.ShopName,
		)
		if err != nil {
			p.log.Error(fmt.Sprintf("err occured while scan rows for GetProductHistoryById: %s", err.Error()))
			return nil, err
		}
		res.ProductHistory = append(res.ProductHistory, &productHistory)
	}

	countQuery := `
		SELECT 
			count(*)
		FROM 
			"product_history" AS p
		WHERE p.product_id = $1
	`
	err = p.db.QueryRow(countQuery, req.ProductId).Scan(&count)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning total GetProductHistory by product_id")
	}

	res.Count = count
	return res, nil
}

func (p *productRepo) GetOrderedProducts(ctx context.Context, req *order_service.CreateOrderCopyRequest) (res []*common.CreateOrderItemCopyRequest, err error) {

	getOrderedProductByOrderIdForImported := `
	SELECT
		oi.product_id,
		pm.quantity,
		ii.name,
		ii.supply_price,
		ii.sale_price
	FROM
		"order" AS o
	JOIN "order_item" AS oi ON o.id = oi.order_id
	LEFT JOIN "product_movement" AS pm ON pm.order_item_id = oi.id
	JOIN "import_item" AS ii ON pm.import_item_id = ii.id
	WHERE
		o.id = $1`
	rows, err := p.db.Query(getOrderedProductByOrderIdForImported, req.GetOrderId())
	if err != nil {
		p.log.Error(fmt.Sprintf("err occured while get rows for getOrderedProductByOrderIdForImported: %s", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			ordererProduct common.CreateOrderItemCopyRequest
		)
		err = rows.Scan(
			&ordererProduct.ProductId,
			&ordererProduct.Value,
			&ordererProduct.Name,
			&ordererProduct.SupplyPrice,
			&ordererProduct.SalePrice,
		)
		if err != nil {
			p.log.Error(fmt.Sprintf("err occured while scan rows for getOrderedProductByOrderIdForImported: %s", err.Error()))
			return nil, err
		}
		res = append(res, &ordererProduct)
	}
	return res, nil
}
