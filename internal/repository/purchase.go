package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/doug-martin/goqu/v9"
)

type PurchaseRepository interface {
	Save(ctx context.Context, purchase *domain.Purchase) (*domain.Purchase, error)
	Update(ctx context.Context, purchase goqu.Record) error
	FindAllWithFilter(ctx context.Context) ([]domain.Purchase, error)
	FindById(ctx context.Context, id string) (domain.Purchase, error)
	Delete(ctx context.Context, id string) error
}

type purchaseRepository struct {
	db *goqu.Database
}

func NewPurchase(db *sql.DB) PurchaseRepository {
	return &purchaseRepository{
		db: goqu.New("default", db),
	}
}

func (d purchaseRepository) Save(ctx context.Context, purchase *domain.Purchase) (*domain.Purchase, error) {
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
