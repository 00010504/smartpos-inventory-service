package postgres

import (
	"context"
	"database/sql"
	"genproto/common"
	"genproto/inventory_service"
	"strings"

	"github.com/Invan2/invan_inventory_service/models"
	"github.com/Invan2/invan_inventory_service/pkg/helper"
	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type supplierRepo struct {
	db  models.DB
	log logger.Logger
}

func NewSupplierRepo(db models.DB, log logger.Logger) repo.SupplierI {
	return &supplierRepo{
		db:  db,
		log: log,
	}
}

func (s *supplierRepo) Create(ctx context.Context, req *inventory_service.CreateSupplierRequest) (*common.ResponseID, error) {

	var (
		supplierId = uuid.NewString()
	)

	query := `
		INSERT INTO "supplier"
			(	"id",
			 	"name",
				"phone_number",
				"company_id",
			    "supplier_company_name",
				"agreement_number", 
				"legal_address", 
				"bank_account", 
				"bank_name", 
				"tin", 
				"ibt", 
				"supplier_company_legal_name"
			)
		VALUES 
			(
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
			)
	`

	_, err := s.db.Exec(query, supplierId, req.Name, req.PhoneNumber, req.Request.CompanyId, req.SupplierCompanyName, req.AgreementNumber, req.LegalAddress, req.BankAccount, req.BankName, req.Tin, req.Ibt, req.SupplierCompanyLegalName)
	if err != nil {
		return nil, errors.Wrap(err, "error while create supplier")
	}

	err = s.insertSupplierFiles(supplierId, req.Files)
	if err != nil {
		return nil, err
	}

	return &common.ResponseID{Id: supplierId}, nil
}

func (s *supplierRepo) GetAll(ctx context.Context, req *common.SearchRequest) (*inventory_service.GetAllSuppliersResponse, error) {
	var (
		res = inventory_service.GetAllSuppliersResponse{
			Data:  make([]*inventory_service.GetSupplierResponse, 0),
			Total: 0,
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
			s.id,
			s.supplier_company_name,
			s.agreement_number,
			s.name,
			s.phone_number,
			s.external_id
		FROM "supplier" s
	`
	filter := ` WHERE s.company_id=:company_id AND s.deleted_at = 0 `
	if req.Search != "" {
		filter += `  AND (
		s."name" ILIKE '%' || :search || '%'
		OR s."agreement_number" ILIKE '%' || :search || '%'
		OR s."phone_number" ILIKE '%' || :search || '%'
		OR s."supplier_company_name" ILIKE '%' || :search || '%'
		) 
		`
	}

	query += filter + `
		ORDER BY  s.created_at DESC
		LIMIT :limit
		OFFSET :offset
	`

	rows, err := s.db.NamedQuery(query, values)
	if err != nil {
		return nil, errors.Wrap(err, "error while search")
	}

	defer rows.Close()

	for rows.Next() {

		var (
			supplier = inventory_service.GetSupplierResponse{}
		)

		err = rows.Scan(&supplier.Id, &supplier.SupplierCompanyName, &supplier.AgreementNumber, &supplier.Name, &supplier.PhoneNumber, &supplier.ExternalId)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning suppliers")
		}

		res.Data = append(res.Data, &supplier)
	}

	query = `
		SELECT
			count(s.id)
		from "supplier" s

	` + filter

	stmt, err := s.db.PrepareNamed(query)
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

func (s *supplierRepo) GetById(ctx context.Context, req *common.RequestID) (*inventory_service.GetSupplierByIdResponse, error) {
	var res inventory_service.GetSupplierByIdResponse

	query := `
	SELECT 
		s.id,
		s.name,
		s.phone_number,
		s.supplier_company_name,
		s.agreement_number,
		s.legal_address,
		s.bank_account,
		s.bank_name,
		s.tin,
		s.ibt,
		s.supplier_company_legal_name
	FROM "supplier" s
	WHERE s.id = $1 AND s.deleted_at = 0
	`

	err := s.db.QueryRow(query, req.Id).Scan(
		&res.Id,
		&res.Name,
		&res.PhoneNumber,
		&res.SupplierCompanyName,
		&res.AgreementNumber,
		&res.LegalAddress,
		&res.BankAccount,
		&res.BankName,
		&res.Tin,
		&res.Ibt,
		&res.SupplierCompanyLegalName)
	if err != nil {
		return nil, errors.Wrap(err, "error while scanning supplier")
	}

	query = `
	SELECT
		sf.file_name,
		sf.name
	FROM "supplier_file" sf
	WHERE sf.supplier_id = $1 AND deleted_at = 0
	`

	rows, err := s.db.Query(query, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "error while get supplier file")
	}

	for rows.Next() {

		var file inventory_service.ProductFile

		err = rows.Scan(&file.FileUrl, &file.Name)
		if err != nil {
			return nil, errors.Wrap(err, "error while scan supplier file")
		}

		res.Files = append(res.Files, &file)
	}

	return &res, nil
}

func (s *supplierRepo) Update(ctx context.Context, req *inventory_service.UpdateSupplierRequest) (*common.ResponseID, error) {

	query := `
		UPDATE
			"supplier"
		SET
			name = $2,
			phone_number = $3,
			agreement_number = $5,
			legal_address = $6,
			bank_account = $7,
			bank_name = $8,
			tin = $9,
			ibt = $10,
			supplier_company_name = $11,
			supplier_company_legal_name = $12
		WHERE
			id = $1 AND deleted_at = 0 AND company_id = $4
	`
	res, err := s.db.Exec(
		query,
		req.Id,
		req.Name,
		req.PhoneNumber,
		req.Request.CompanyId,
		req.AgreementNumber,
		req.LegalAddress,
		req.BankAccount,
		req.BankName,
		req.Tin,
		req.Ibt,
		req.SupplierCompanyName,
		req.SupplierCompanyLegalName,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error while update supplier")

	}
	if i, _ := res.RowsAffected(); i == 0 {
		return nil, sql.ErrNoRows
	}

	if err = s.deleteSupplierFiles(req.Id); err != nil {
		return nil, err
	}

	if err = s.insertSupplierFiles(req.Id, req.Files); err != nil {
		return nil, err
	}

	return &common.ResponseID{Id: req.Id}, nil
}

func (s *supplierRepo) insertSupplierFiles(supplierId string, files []*inventory_service.ProductFile) error {

	values := []interface{}{}

	if len(files) == 0 {
		return nil
	}

	query := `
		INSERT INTO
			"supplier_file"
		(
			id,
			supplier_id,
			file_name,
			name
		)
		VALUES 
	`

	for _, file := range files {
		res := strings.Split(file.FileUrl, "/")

		query += "(?, ?, ?, ?),"
		values = append(values,
			uuid.New().String(),
			supplierId,
			res[len(res)-1],
			file.Name,
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return errors.Wrap(err, "error while insert supplier_file. Prepare")
	}

	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return errors.Wrap(err, "error while insert supplier_file. Exec")
	}

	return nil
}

func (s *supplierRepo) deleteSupplierFiles(supplierId string) error {

	query := `
		UPDATE
			"supplier_file"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE
			supplier_id = $1 AND deleted_at = 0
	`

	_, err := s.db.Exec(query, supplierId)
	if err != nil {
		return errors.Wrap(err, "error while delete file")
	}

	return nil
}

func (s *supplierRepo) Delete(ctx context.Context, req *common.RequestID) (*common.ResponseID, error) {
	query := `
		UPDATE 
			"supplier"
		SET
			deleted_at = extract(epoch from now())::bigint
		WHERE 
			id = $1 AND deleted_at = 0
		`
	res, err := s.db.Exec(query, req.Id)
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
