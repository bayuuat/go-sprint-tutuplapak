package service

import (
	"context"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type purchaseService struct {
	cnf                *config.Config
	purchaseRepository repository.PurchaseRepository
}

type PurchaseService interface {
	GetPurchasesWithFilter(ctx context.Context, filter dto.PurchaseFilter, userId string) ([]dto.PurchaseData, int, error)
	CreatePurchase(ctx context.Context, req dto.PurchaseReq, userId string) (dto.PurchaseData, int, error)
	PatchPurchase(ctx context.Context, req dto.UpdatePurchaseReq, userId, id string) (dto.PurchaseData, int, error)
	DeletePurchase(ctx context.Context, user_id, id string) (dto.PurchaseData, int, error)
}

func NewPurchase(cnf *config.Config,
	purchaseRepository repository.PurchaseRepository) PurchaseService {
	return &purchaseService{
		cnf:                cnf,
		purchaseRepository: purchaseRepository,
	}
}

func (ds purchaseService) GetPurchasesWithFilter(ctx context.Context, filter dto.PurchaseFilter, userId string) ([]dto.PurchaseData, int, error) {
	return []dto.PurchaseData{}, 200, nil
}

func (ds *purchaseService) CreatePurchase(ctx context.Context, req dto.PurchaseReq, userId string) (dto.PurchaseData, int, error) {
	return dto.PurchaseData{}, 200, nil
}

func (ds purchaseService) PatchPurchase(ctx context.Context, req dto.UpdatePurchaseReq, userId, id string) (dto.PurchaseData, int, error) {
	return dto.PurchaseData{}, 200, nil
}

func (ds purchaseService) DeletePurchase(ctx context.Context, userId, id string) (dto.PurchaseData, int, error) {
	return dto.PurchaseData{}, 200, nil
}
