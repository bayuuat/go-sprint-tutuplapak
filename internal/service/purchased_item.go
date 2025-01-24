package service

import (
	"context"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type purchasedItemService struct {
	cnf                     *config.Config
	purchasedItemRepository repository.PurchasedItemRepository
}

type PurchasedItemService interface {
	GetPurchasedItemsWithFilter(ctx context.Context, filter dto.PurchasedItemFilter, userId string) ([]dto.PurchasedItemData, int, error)
	CreatePurchasedItem(ctx context.Context, req dto.PurchasedItemReq, userId string) (dto.PurchasedItemData, int, error)
	PatchPurchasedItem(ctx context.Context, req dto.UpdatePurchasedItemReq, userId, id string) (dto.PurchasedItemData, int, error)
	DeletePurchasedItem(ctx context.Context, user_id, id string) (dto.PurchasedItemData, int, error)
}

func NewPurchasedItem(cnf *config.Config,
	purchasedItemRepository repository.PurchasedItemRepository,
	purchasedItemTypesRepository repository.PurchasedItemTypesRepository) PurchasedItemService {
	return &purchasedItemService{
		cnf:                     cnf,
		purchasedItemRepository: purchasedItemRepository,
	}
}

func (ds purchasedItemService) GetPurchasedItemsWithFilter(ctx context.Context, filter dto.PurchasedItemFilter, userId string) ([]dto.PurchasedItemData, int, error) {
	return nil, nil, nil
}

func (ds *purchasedItemService) CreatePurchasedItem(ctx context.Context, req dto.PurchasedItemReq, userId string) (dto.PurchasedItemData, int, error) {
	return nil, nil, nil
}

func (ds purchasedItemService) PatchPurchasedItem(ctx context.Context, req dto.UpdatePurchasedItemReq, userId, id string) (dto.PurchasedItemData, int, error) {
	return nil, nil, nil
}

func (ds purchasedItemService) DeletePurchasedItem(ctx context.Context, userId, id string) (dto.PurchasedItemData, int, error) {
	return nil, nil, nil
}
