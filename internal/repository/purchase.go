package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type PurchaseRepository interface {
	Save(ctx context.Context, purchase *domain.Purchase) (*domain.Purchase, error)
	Update(ctx context.Context, userId string, purchase goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.PurchaseFilter, userId string) ([]domain.Purchase, error)
	FindById(ctx context.Context, userId, id string) (domain.Purchase, error)
	Delete(ctx context.Context, userId, id string) error
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
	return nil, nil
}

func (d purchaseRepository) Update(ctx context.Context, userId string, purchase goqu.Record) error {
	return nil
}

func (d purchaseRepository) FindById(ctx context.Context, userId, id string) (purchase domain.Purchase, err error) {
	return nil, nil
}

func (d purchaseRepository) Delete(ctx context.Context, userId, id string) error {
	return nil
}

func (d purchaseRepository) FindAllWithFilter(ctx context.Context, filter *dto.PurchaseFilter, userId string) ([]domain.Purchase, error) {
	return nil, nil
}
