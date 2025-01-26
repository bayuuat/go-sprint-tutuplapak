package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type PurchasedItemRepository interface {
	Save(ctx context.Context, purchasedItem *domain.PurchasedItem) (*domain.PurchasedItem, error)
	Update(ctx context.Context, userId string, purchasedItem goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.PurchasedItemFilter, userId string) ([]domain.PurchasedItem, error)
	FindById(ctx context.Context, userId, id string) (domain.PurchasedItem, error)
	Delete(ctx context.Context, userId, id string) error
}

type purchasedItemRepository struct {
	db *goqu.Database
}

func NewPurchasedItem(db *sql.DB) PurchasedItemRepository {
	return &purchasedItemRepository{
		db: goqu.New("default", db),
	}
}

func (d purchasedItemRepository) Save(ctx context.Context, purchasedItem *domain.PurchasedItem) (*domain.PurchasedItem, error) {
	return nil, nil
}

func (d purchasedItemRepository) Update(ctx context.Context, userId string, purchasedItem goqu.Record) error {
	return nil
}

func (d purchasedItemRepository) FindById(ctx context.Context, userId, id string) (purchasedItem domain.PurchasedItem, err error) {
	return nil, nil
}

func (d purchasedItemRepository) Delete(ctx context.Context, userId, id string) error {
	return nil
}

func (d purchasedItemRepository) FindAllWithFilter(ctx context.Context, filter *dto.PurchasedItemFilter, userId string) ([]domain.PurchasedItem, error) {
	return nil, nil
}
