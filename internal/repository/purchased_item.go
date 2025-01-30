package repository

import (
	"context"
	"errors"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type PurchasedItemRepository interface {
	Save(ctx context.Context, purchasedItem *domain.PurchasedItem) (*domain.PurchasedItem, error)
	SavesTx(ctx context.Context, tx *goqu.TxDatabase, purchasedItems []domain.PurchasedItemReq) error
	Update(ctx context.Context, userId string, purchasedItem goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.PurchasedItemFilter, userId string) ([]domain.PurchasedItem, error)
	FindById(ctx context.Context, userId, id string) (domain.PurchasedItem, error)
	Delete(ctx context.Context, userId, id string) error
}

type purchasedItemRepository struct {
	db *goqu.Database
}

func NewPurchasedItem(db *goqu.Database) PurchasedItemRepository {
	return &purchasedItemRepository{
		db: db,
	}
}

func (d purchasedItemRepository) Save(ctx context.Context, purchasedItem *domain.PurchasedItem) (*domain.PurchasedItem, error) {
	return &domain.PurchasedItem{}, errors.New("not implemented")
}

func (d purchasedItemRepository) SavesTx(ctx context.Context, tx *goqu.TxDatabase, purchasedItems []domain.PurchasedItemReq) error {
	_, err := tx.Insert("purchased_items").Rows(purchasedItems).Executor().ExecContext(ctx)
	return err
}

func (d purchasedItemRepository) Update(ctx context.Context, userId string, purchasedItem goqu.Record) error {
	return errors.New("not implemented")
}

func (d purchasedItemRepository) FindById(ctx context.Context, userId, id string) (purchasedItem domain.PurchasedItem, err error) {
	return domain.PurchasedItem{}, errors.New("not implemented")
}

func (d purchasedItemRepository) Delete(ctx context.Context, userId, id string) error {
	return errors.New("not implemented")
}

func (d purchasedItemRepository) FindAllWithFilter(ctx context.Context, filter *dto.PurchasedItemFilter, userId string) ([]domain.PurchasedItem, error) {
	return []domain.PurchasedItem{}, errors.New("not implemented")
}
