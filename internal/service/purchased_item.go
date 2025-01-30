package service

import (
	"context"
	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"github.com/doug-martin/goqu/v9"
	"net/http"
)

type PurchasedItemService struct {
	cnf                     *config.Config
	purchasedItemRepository repository.PurchasedItemRepository
	ProductService          ProductServicer
	FileService             FileServicer
}

type PurchasedItemServicer interface {
	CreatePurchasedItemsTx(ctx context.Context, tx *goqu.TxDatabase, req []dto.PurchasedItemReq) ([]dto.ProductData, map[string]int, int, error)
}

func NewPurchasedItemServicer(cnf *config.Config,
	purchasedItemRepository repository.PurchasedItemRepository,
	productService ProductService,
	fileService FileService,
) PurchasedItemServicer {
	return &PurchasedItemService{
		cnf:                     cnf,
		purchasedItemRepository: purchasedItemRepository,
		ProductService:          productService,
		FileService:             fileService,
	}
}

func (ds *PurchasedItemService) CreatePurchasedItemsTx(ctx context.Context, tx *goqu.TxDatabase, req []dto.PurchasedItemReq) (
	purchasedItems []dto.ProductData, totalPerSeller map[string]int, code int, err error) {
	productIds := make([]string, len(req))
	for i, purchasedItem := range req {
		productIds[i] = purchasedItem.ProductID
	}

	products, code, err := ds.ProductService.GetProducts(ctx, productIds)
	if err != nil {
		return nil, nil, code, err
	}

	// TODO: Check stock?

	price := make(map[string]int)
	for _, product := range products {
		price[product.ProductId] = product.Price
	}

	productDbInsert := make([]domain.PurchasedItemReq, len(req))
	for _, productReq := range req {
		productDbInsert = append(productDbInsert, domain.PurchasedItemReq{
			PurchaseID: productReq.PurchaseID,
			ProductID:  productReq.ProductID,
			Qty:        productReq.Qty,
			Price:      price[productReq.ProductID],
		})
	}

	err = ds.purchasedItemRepository.SavesTx(ctx, tx, productDbInsert)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return products, nil, http.StatusCreated, nil
}
