package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"net/http"
	"strconv"
)

type PurchaseService struct {
	cnf                  *config.Config
	purchaseRepository   repository.PurchaseRepository
	purchasedItemService PurchasedItemService
	userService          userService
}

type PurchaseServicer interface {
	CreatePurchase(ctx context.Context, req *dto.PurchaseReq) (dto.PurchaseData, int, error)
	CreatePayment(ctx context.Context, req dto.PurchasedItemReq) (int, error)
}

func NewPurchaseServicer(cnf *config.Config,
	purchaseRepository repository.PurchaseRepository,
	purchasedItemService PurchasedItemService,
	userService userService,
) PurchaseServicer {
	return PurchaseService{
		cnf:                  cnf,
		purchaseRepository:   purchaseRepository,
		purchasedItemService: purchasedItemService,
		userService:          userService,
	}
}

func (ds PurchaseService) CreatePurchase(ctx context.Context, req *dto.PurchaseReq) (dto.PurchaseData, int, error) {
	tx, err := ds.purchaseRepository.BeginTx()
	if err != nil {
		return dto.PurchaseData{}, http.StatusInternalServerError, err
	}

	purchaseId, err := ds.purchaseRepository.SaveTx(ctx, tx, domain.PurchaseReq{
		SenderName:          req.SenderName,
		SenderContactType:   req.SenderContactType,
		SenderContactDetail: req.SenderContactDetail,
	})

	if err != nil {
		err2 := ds.purchaseRepository.RollbackTx(tx)
		if err2 != nil {
			err = fmt.Errorf("failed to rollback transaction: %w; %w", err2, err)
		}
		return dto.PurchaseData{}, http.StatusInternalServerError, err
	}

	for _, purchaseItem := range req.PurchasedItemsReq {
		purchaseItem.PurchaseID = strconv.Itoa(purchaseId)
	}

	purchasedItems, totalPerSeller, code, err := ds.purchasedItemService.CreatePurchasedItemsTx(ctx, tx, req.PurchasedItemsReq)
	if err != nil {
		err2 := ds.purchaseRepository.RollbackTx(tx)
		if err2 != nil {
			err = fmt.Errorf("failed to rollback transaction: %w; %w", err2, err)
		}
		return dto.PurchaseData{}, code, err
	}

	sellerIds := getSellerIds(totalPerSeller)
	userFilter := []string{"id", "bank_account_name", "bank_account_holder", "bank_account_number"}
	sellersData, code, err := ds.userService.GetUsersFilter(ctx, sellerIds, userFilter)
	if err != nil {
		return dto.PurchaseData{}, http.StatusInternalServerError, err
	}

	paymentDetails := make([]dto.PaymentDetail, len(sellersData))
	for i, seller := range sellersData {
		paymentDetails[i] = dto.PaymentDetail{
			PaymentDetailData: sellersData[i],
			TotalPrice:        totalPerSeller[seller.SellerID],
		}
	}

	return dto.PurchaseData{
		PurchaseID:     strconv.Itoa(purchaseId),
		PurchasedItems: purchasedItems,
		PaymentDetails: paymentDetails,
	}, http.StatusCreated, nil
}

func (ds PurchaseService) CreatePayment(ctx context.Context, req dto.PurchasedItemReq) (int, error) {
	return 0, errors.New("not implemented")
}

func getSellerIds(totalPerSeller map[string]int) (sellerIds []string) {
	sellerIds = make([]string, len(totalPerSeller))
	i := 0
	for k := range totalPerSeller {
		sellerIds[i] = k
		i++
	}
	return
}
