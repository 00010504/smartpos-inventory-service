package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"genproto/common"
	"genproto/inventory_service"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type supplierOrderRepo struct {
	db  models.DB
	log logger.Logger
}

func NewSupplierOrderRepo(db models.DB, log logger.Logger) repo.SupplierOrderI {
	return &supplierOrderRepo{
		db:  db,
		log: log,
	}
}

func (sp *supplierOrderRepo) Create(ctx context.Context, req *inventory_service.CreateSupplierOrderRequest) (*common.ResponseID, error) {

	var (
		supplierOrderId = uuid.NewString()
	)

	query := `
	INSERT INTO "supplier_order"
		(
			"id",
			"supplier_id",
			"invoice_number",
			"shop_id",
			"company_id",
			"expected_date",
			"note",
			"created_by"
		)
	VALUES 
		(
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`
	_, err := sp.db.Exec(
		query,
		supplierOrderId,
		req.SupplierId,
		req.InvoiceNumber,
		req.ShopId,
		req.Request.CompanyId,
		req.ExpectedDate,
		req.Note,
		req.Request.UserId,
	)

	if err != nil {
		return nil, errors.Wrap(err, "error while create supplier order")
	}

	return &common.ResponseID{Id: supplierOrderId}, nil
}

func (sp *supplierOrderRepo) UpsertOrderItem(ctx context.Context, req *inventory_service.CreateSupplierOrderItemRequest) (string, error) {

	var (
		supplierOrderItemId = uuid.NewString()
	)

	query := `
			INSERT INTO
				"supplier_order_item"
			(
				id,
				supplier_order_id,
				product_id,
				expected_amount,
				discount,
				cost,
				sold_amount
			)
			VALUES
			(
				$1, $2, $3, $4, $5, $6, $7
			) ON CONFLICT (supplier_order_id, product_id, deleted_at) DO
			UPDATE
				SET
					supplier_order_id = $2,
					product_id = $3,
					expected_amount = $4,
					discount = $5,
					cost = $6,
					sold_amount = $7
			RETURNING id
		`

	err := sp.db.QueryRow(
		query,
		supplierOrderItemId,
		req.SupplierOrderId,
		req.ProductId,
		req.ExpectedAmount,
		req.Discount,
		req.Cost,
		req.SoldAmount,
	).Scan(&supplierOrderItemId)
	if err != nil {
		return "", errors.Wrap(err, "error while create supplier order item")
	}

	return supplierOrderItemId, nil
}

func (sp *supplierOrderRepo) GetOrderItem(supplierOrderItemId string) (*inventory_service.CreateSupplierOrderItemElasticResponse, error) {

	var res = inventory_service.CreateSupplierOrderItemElasticResponse{
		Barcode: make([]string, 0),
	}

	query := `
		SELECT
			soi.id,
			soi.product_id,
			p.name,
			soi.expected_amount,
			soi.discount,
			soi.cost,
			soi.sold_amount
		FROM "supplier_order_item" soi
		LEFT JOIN "product" p ON soi.product_id = p.id AND p.deleted_at = 0
		WHERE soi.id = $1 AND soi.deleted_at = 0
	`

	err := sp.db.QueryRow(query, supplierOrderItemId).Scan(
		&res.Id,
		&res.ProductId,
		&res.ProductName,
		&res.ExpectedAmount,
		&res.Discount,
		&res.Cost,
		&res.SoldAmount,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning supplier order item")
	}

	productBarcodes, err := sp.getOrderItemBarcodes([]string{res.ProductId})
	if err != nil {
		return nil, errors.Wrap(err, "error while scan supplier order item")
	}

	res.Barcode = productBarcodes[res.ProductId]

	return &res, nil
}

func (sp *supplierOrderRepo) GetAll(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSupplierOrderResponse, error) {
	var (
		res = inventory_service.GetAllSupplierOrderResponse{
			Data: make([]*inventory_service.GetSupplierOrdersResponse, 0),
		}
		values = map[string]interface{}{
			"limit":      req.Limit,
			"offset":     req.Limit * (req.Page - 1),
			"search":     req.Search,
			"company_id": req.Request.CompanyId,
		}
	)

	query := `
		SELECT
			so.id,
			s.id,
			s.name,
			so.external_id,
			u.id,
			u.first_name,
			u.last_name,
			shp.id,
			shp.name,
			sos.id,
			sos.name,
			sos.translation,
			so.expected_date,
			so.total_sold_amount

		FROM "supplier_order" so
		LEFT JOIN "user" u ON u.id = so.created_by AND u.deleted_at = 0
		LEFT JOIN "shop" shp ON shp.id = so.shop_id AND shp.deleted_at = 0
		JOIN "supplier_order_status" sos ON sos.id = so.status_id
		JOIN "supplier" s ON s.id = so.supplier_id AND s.deleted_at = 0
	`

	filter := ` WHERE so.company_id=:company_id AND so.deleted_at = 0 `
	if req.Search != "" {
		filter += `  AND (
		so."external_id" ILIKE '%' || :search || '%'
		) 
		`
	}

	query += filter + `
		ORDER BY  so.created_at DESC
		LIMIT :limit
		OFFSET :offset

	`

	rows, err := sp.db.NamedQuery(query, values)
	if err != nil {
		return nil, errors.Wrap(err, "error while search")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			supplierO = inventory_service.GetSupplierOrdersResponse{
				SupplierName: &inventory_service.SupplierOrderName{},
				Status:       &inventory_service.SupplierOrderStatus{},
			}
			shortUser models.NullShortUser
			shortShop models.NullShortShop
			tr        []byte
		)

		err = rows.Scan(
			&supplierO.Id,
			&supplierO.SupplierName.Id,
			&supplierO.SupplierName.Name,
			&supplierO.ExternalId,
			&shortUser.ID,
			&shortUser.FirstName,
			&shortUser.LastName,
			&shortShop.ID,
			&shortShop.Name,
			&supplierO.Status.Id,
			&supplierO.Status.Name,
			&tr,
			&supplierO.ExpectedDate,
			&supplierO.TotalAmount,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning supplier order")
		}

		if err = json.Unmarshal(tr, &supplierO.Status.Translation); err != nil {
			return nil, errors.Wrap(err, "error while Unmarshal StatusTranslation")
		}

		supplierO.Shop = &common.ShortShop{
			Id:   shortShop.ID.String,
			Name: shortShop.Name.String,
		}

		supplierO.User = &common.ShortUser{
			Id:        shortUser.ID.String,
			FirstName: shortUser.FirstName.String,
			LastName:  shortUser.LastName.String,
		}

		res.Data = append(res.Data, &supplierO)
	}
	query = `
		SELECT
			count(so.id)
		FROM "supplier_order" so

	` + filter

	stmt, err := sp.db.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrap(err, "error while prepareName")
	}

	defer stmt.Close()

	err = stmt.QueryRow(values).Scan(&res.Total)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning queryRow")
	}

	return &res, nil
}

func (sp *supplierOrderRepo) GetById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierOrderByIdResponse, error) {

	var (
		res = inventory_service.GetSupplierOrderByIdResponse{
			Status:       &inventory_service.SupplierOrderStatus{},
			SupplierName: &inventory_service.SupplierOrderName{},
		}
		shortShop   models.NullShortShop
		tr          []byte
		product_ids = make([]string, 0)
	)

	query := `
		SELECT 
			so.id,
			sos.id,
			sos.name,
			sos.translation,
			s.id,
			s.name,
			so.invoice_number,
			shp.id,
			shp.name,
			so.expected_date,
			so.created_at,
			so.note
		FROM "supplier_order" so
		LEFT JOIN "shop" shp ON shp.id = so.shop_id AND shp.deleted_at = 0
		JOIN "supplier_order_status" sos ON sos.id = so.status_id
		JOIN "supplier" s ON s.id = so.supplier_id AND s.deleted_at = 0
		WHERE so.id = $1 AND so.deleted_at = 0
	`

	err := sp.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Status.Id,
		&res.Status.Name,
		&tr,
		&res.SupplierName.Id,
		&res.SupplierName.Name,
		&res.InvoiceNumber,
		&shortShop.ID,
		&shortShop.Name,
		&res.ExpectedDate,
		&res.CreatedDate,
		&res.Note,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning supplier order")
	}
	if err = json.Unmarshal(tr, &res.Status.Translation); err != nil {
		return nil, errors.Wrap(err, "error while Unmarshal StatusTranslation")
	}

	res.Shop = &common.ShortShop{
		Id:   shortShop.ID.String,
		Name: shortShop.Name.String,
	}

	query = `
	SELECT
		soi.id,
		soi.supplier_order_id,
		soi.product_id,
		p.name,
		soi.expected_amount,
		soi.discount,
		soi.sold_amount,
		soi.cost
	FROM "supplier_order_item" soi
	LEFT JOIN "product" p ON p.id = soi.product_id AND p.deleted_at = 0
	WHERE soi.supplier_order_id = $1 AND soi.deleted_at = 0
	`

	rows, err := sp.db.Query(query, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error while get supplier order item")
	}

	for rows.Next() {

		var item inventory_service.GetSupplierOrderItemByIdResponse

		err = rows.Scan(&item.Id,
			&item.SupplierOrderId,
			&item.ProductId,
			&item.ProductName,
			&item.ExpectedAmount,
			&item.Discount,
			&item.SoldAmount,
			&item.Cost)
		if err != nil {
			return nil, errors.Wrap(err, "error while scan supplier order item")
		}

		product_ids = append(product_ids, item.ProductId)
		res.Items = append(res.Items, &item)
	}

	productBarcodes, err := sp.getOrderItemBarcodes(product_ids)
	if err != nil {
		return nil, errors.Wrap(err, "error while scan supplier order item")
	}

	for _, item := range res.Items {
		item.Barcode = productBarcodes[item.ProductId]
	}

	return &res, nil
}

func (sp *supplierOrderRepo) getOrderItemBarcodes(ids []string) (map[string][]string, error) {

	var res = make(map[string][]string, 0)

	for _, id := range ids {
		res[id] = make([]string, 0)
	}

	query := `
		SELECT "barcode", "product_id"
		FROM "product_barcode"
		WHERE product_id = ANY($1)
	`

	rows, err := sp.db.Query(query, pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "error while getting supplier order item product_barcode. Query")
	}

	defer rows.Close()

	for rows.Next() {
		var (
			barcode    string
			product_id string
		)

		err = rows.Scan(&barcode, &product_id)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning supplier order item product_barcode")
		}

		res[product_id] = append(res[product_id], barcode)
	}

	return res, nil
}

func (sp *supplierOrderRepo) UpdateSupplierOrderStatus(ctx context.Context, req *common.RequestID, supplierOrderStatus string) (*common.ResponseID, error) {
	query := `
		UPDATE
			"supplier_order"
		SET
			status_id = $2
		WHERE 
			id = $1 AND deleted_at = 0
	`
	res, err := sp.db.Exec(
		query,
		req.Id,
		supplierOrderStatus,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while update supplier order status")

	}
	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	return &common.ResponseID{Id: req.Id}, nil
}

func (i *supplierOrderRepo) FinishSupplierOrder(ctx context.Context, req *inventory_service.FinishSupplierOrderRequest) error {
	query := `update supplier_order set
		status_id = $2
		WHERE id = $1 `
	_, err := i.db.Exec(
		query,
		req.SupplierOrderId,
		"caba1cc0-ba0d-4a03-8b78-c81c6730cca6",
	)
	if err != nil {
		return errors.Wrap(err, "error while finish supplier_order")
	}
	return nil
}

func (sp *supplierOrderRepo) UpdateSupplierOrderAmount(ctx context.Context, req *inventory_service.UpdateSupplierOrderRecivedRequest) (*common.ResponseID, error) {

	query := `
		UPDATE
			"supplier_order_item"
		SET
			arrived_amount = $2
		WHERE
			id = $1 AND deleted_at = 0
	`
	_, err := sp.db.Exec(
		query,
		req.Id,
		req.ArrivedAmount,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while update supplier order item arrived amount")
	}

	return &common.ResponseID{Id: req.Id}, nil
}

func (sp *supplierOrderRepo) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {

	query := `
		UPDATE 
			"supplier_order"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE 
			id = $1 AND deleted_at = 0
		`
	res, err := sp.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if i == 0 {
		return nil, sql.ErrNoRows
	}

	return &common.ResponseID{
		Id: req.Id,
	}, nil
}

func (sp *supplierOrderRepo) DeleteItemById(ctx context.Context, req *common.RequestID) error {
	query := `
		UPDATE 
			"supplier_order_item"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE 
			id = $1 AND deleted_at = 0
		`
	res, err := sp.db.Exec(query, req.Id)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
