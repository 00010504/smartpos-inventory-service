package storage

import (
	"context"
	"database/sql"

	"github.com/Invan2/invan_inventory_service/pkg/logger"
	"github.com/Invan2/invan_inventory_service/storage/postgres"
	"github.com/Invan2/invan_inventory_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type repos struct {
	companyRepo        repo.CompanyI
	userRepo           repo.UserI
	shopRepo           repo.ShopI
	productRepo        repo.ProductI
	orderRepo          repo.OrderI
	cashboxRepo        repo.CashboxI
	measurementRepo    repo.MeasurementI
	importRepo         repo.ImportI
	supplierRepo       repo.SupplierI
	supplier_orderRepo repo.SupplierOrderI
}

type repoIs interface {
	Company() repo.CompanyI
	User() repo.UserI
	Shop() repo.ShopI
	Product() repo.ProductI
	Order() repo.OrderI
	Cashbox() repo.CashboxI
	Measurement() repo.MeasurementI
	Import() repo.ImportI
	Supplier() repo.SupplierI
	SupplierO() repo.SupplierOrderI
}

type storage struct {
	db  *sqlx.DB
	log logger.Logger
	repos
}

type storageTr struct {
	tr *sqlx.Tx
	repos
}

type StorageTrI interface {
	Commit() error
	Rollback() error
	repoIs
}

type StoragePgI interface {
	Product() repo.ProductI
	WithTransaction(ctx context.Context) (StorageTrI, error)
	repoIs
}

func NewStoragePg(ctx context.Context, log logger.Logger, db *sqlx.DB) (StoragePgI, error) {

	return &storage{
		db:  db,
		log: log,
		repos: repos{
			companyRepo:        postgres.NewCompanyRepo(log, db),
			userRepo:           postgres.NewUserRepo(log, db),
			shopRepo:           postgres.NewShopRepo(db, log),
			productRepo:        postgres.NewProductRepo(log, db),
			orderRepo:          postgres.NewOrderRepo(db, log),
			cashboxRepo:        postgres.NewCashboxRepo(db, log),
			measurementRepo:    postgres.NewMeasurementRepo(db, log),
			importRepo:         postgres.NewImportRepo(db, log),
			supplierRepo:       postgres.NewSupplierRepo(db, log),
			supplier_orderRepo: postgres.NewSupplierOrderRepo(db, log),
		},
	}, nil
}

func (s *storage) WithTransaction(ctx context.Context) (StorageTrI, error) {
	tr, err := s.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &storageTr{
		tr: tr,
		repos: repos{
			companyRepo:        postgres.NewCompanyRepo(s.log, tr),
			userRepo:           postgres.NewUserRepo(s.log, tr),
			shopRepo:           postgres.NewShopRepo(tr, s.log),
			productRepo:        postgres.NewProductRepo(s.log, tr),
			orderRepo:          postgres.NewOrderRepo(tr, s.log),
			cashboxRepo:        postgres.NewCashboxRepo(tr, s.log),
			measurementRepo:    postgres.NewMeasurementRepo(tr, s.log),
			importRepo:         postgres.NewImportRepo(tr, s.log),
			supplierRepo:       postgres.NewSupplierRepo(tr, s.log),
			supplier_orderRepo: postgres.NewSupplierOrderRepo(tr, s.log),
		},
	}, nil
}

func (s *storageTr) Commit() error {
	return s.tr.Commit()
}

func (s *storageTr) Rollback() error {
	return s.tr.Rollback()
}

func (r *repos) Company() repo.CompanyI {
	return r.companyRepo
}

func (r *repos) User() repo.UserI {
	return r.userRepo
}

func (r *repos) Shop() repo.ShopI {
	return r.shopRepo
}

func (r *repos) Product() repo.ProductI {
	return r.productRepo
}

func (r *repos) Order() repo.OrderI {
	return r.orderRepo
}

func (r *repos) Cashbox() repo.CashboxI {
	return r.cashboxRepo
}

func (r *repos) Measurement() repo.MeasurementI {
	return r.measurementRepo
}

func (r *repos) Import() repo.ImportI {
	return r.importRepo
}

func (r *repos) Supplier() repo.SupplierI {
	return r.supplierRepo
}

func (r *repos) SupplierO() repo.SupplierOrderI {
	return r.supplier_orderRepo
}
