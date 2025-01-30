package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bayuuat/tutuplapak/domain"
	"github.com/doug-martin/goqu/v9"
)

type PurchaseRepository interface {
	BeginTx() (*goqu.TxDatabase, error)
	CommitTx(tx *goqu.TxDatabase) error
	SaveTx(ctx context.Context, tx *goqu.TxDatabase, purchase domain.PurchaseReq) (purchaseId int, err error)
	RollbackTx(tx *goqu.TxDatabase) error
	Update(ctx context.Context, purchase goqu.Record) error
	FindAllWithFilter(ctx context.Context) ([]domain.Purchase, error)
	FindById(ctx context.Context, id string) (domain.Purchase, error)
	Delete(ctx context.Context, id string) error
}

type purchaseRepository struct {
	db *goqu.Database
}

func NewPurchase(db *goqu.Database) PurchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

func (d purchaseRepository) BeginTx() (*goqu.TxDatabase, error) {
	tx, err := d.db.Begin()
	return tx, err
}

func (d purchaseRepository) CommitTx(tx *goqu.TxDatabase) error {
	return tx.Commit()
}

func (d purchaseRepository) RollbackTx(tx *goqu.TxDatabase) error {
	return tx.Rollback()
}

func (d purchaseRepository) SaveTx(ctx context.Context, tx *goqu.TxDatabase, purchase domain.PurchaseReq) (purchaseId int, err error) {
	purchaseId = 0
	insert := tx.Insert("purchases").Rows(purchase).Returning("purchase_id").Executor()
	_, err = insert.ScanVal(&purchaseId)
	fmt.Println(err)
	return purchaseId, err
}

func (d purchaseRepository) Save(ctx context.Context) (*domain.Purchase, error) {
	return &domain.Purchase{}, errors.New("not implemented")
}

func (d purchaseRepository) Update(ctx context.Context, purchase goqu.Record) error {
	return errors.New("not implemented")
}

func (d purchaseRepository) FindById(ctx context.Context, id string) (purchase domain.Purchase, err error) {
	return domain.Purchase{}, errors.New("not implemented")
}

func (d purchaseRepository) Delete(ctx context.Context, id string) error {
	return errors.New("not implemented")
}

func (d purchaseRepository) FindAllWithFilter(ctx context.Context) ([]domain.Purchase, error) {
	return []domain.Purchase{}, errors.New("not implemented")
}
