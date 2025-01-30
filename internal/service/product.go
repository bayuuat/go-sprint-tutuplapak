package service

import (
	"context"
	"errors"
	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type ProductService struct {
	cnf               *config.Config
	productRepository repository.ProductRepository
	fileService       FileService
}

type ProductServicer interface {
	GetProducts(ctx context.Context, productIds []string) ([]dto.ProductData, int, error)
	GetProductsWithFilter(ctx context.Context, filter dto.ProductFilter, userId string) ([]dto.ProductData, int, error)
	CreateProduct(ctx context.Context, req dto.ProductReq, userId string) (dto.ProductData, int, error)
	PutProduct(ctx context.Context, req dto.ProductReq, id string) (dto.ProductData, int, error)
	DeleteProduct(ctx context.Context, userId, id string) (dto.ProductData, int, error)
}

func NewProductServicer(cnf *config.Config,
	productRepository repository.ProductRepository,
	fileService FileService) ProductServicer {
	return ProductService{
		cnf:               cnf,
		productRepository: productRepository,
		fileService:       fileService,
	}
}

func (ds ProductService) GetProducts(ctx context.Context, productIds []string) ([]dto.ProductData, int, error) {
	res, err := ds.productRepository.FindByIds(ctx, productIds)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	files, code, err := ds.fileService.GetFileIds(ctx, productIds)
	if err != nil {
		return nil, code, err
	}

	filesMap := make(map[string]dto.FileData, len(files))
	for _, file := range files {
		filesMap[file.FileID] = file
	}

	products := make([]dto.ProductData, len(res))

	for i, product := range res {
		products[i] = dto.ProductData{
			ProductId:        strconv.Itoa(product.ProductID),
			Name:             product.Name,
			Category:         product.Category,
			Qty:              product.Qty,
			Price:            product.Price,
			Sku:              product.SKU,
			FileId:           strconv.Itoa(product.FileID),
			FileUri:          filesMap[strconv.Itoa(product.FileID)].FileID,
			FileThumbnailUri: filesMap[strconv.Itoa(product.FileID)].FileThumbnailURI,
			CreatedAt:        product.CreatedAt.Format(time.RFC3339Nano),
			UpdatedAt:        product.UpdatedAt.Format(time.RFC3339Nano),
		}
	}

	return products, http.StatusOK, nil
}

func (ds ProductService) GetProductsWithFilter(ctx context.Context, filter dto.ProductFilter, userId string) ([]dto.ProductData, int, error) {
	return []dto.ProductData{}, 0, nil
}

func (ds ProductService) CreateProduct(ctx context.Context, req dto.ProductReq, userId string) (dto.ProductData, int, error) {
	return dto.ProductData{}, 0, nil
}

func (ds ProductService) PutProduct(ctx context.Context, req dto.ProductReq, productId string) (dto.ProductData, int, error) {
	file, product, errorCode, err := ds.validateProduct(ctx, req, productId)
	if err != nil {
		return dto.ProductData{}, errorCode, err
	}

	err = ds.productRepository.Put(ctx, dto.Product(req))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.ProductData{}, http.StatusInternalServerError, err
	}

	return dto.ProductData{
		ProductId:        strconv.Itoa(product.ProductID),
		Name:             req.Name,
		Category:         req.Category,
		Qty:              req.Qty,
		Price:            req.Price,
		Sku:              req.Sku,
		FileId:           file.FileID,
		FileUri:          file.FileURI,
		FileThumbnailUri: file.FileThumbnailURI,
		CreatedAt:        time.Now().Format(time.RFC3339Nano),
		UpdatedAt:        time.Now().Format(time.RFC3339Nano),
	}, 200, nil
}

func (ds ProductService) DeleteProduct(ctx context.Context, userId, id string) (dto.ProductData, int, error) {
	return dto.ProductData{}, 0, nil
}

func (ds ProductService) validateProduct(ctx context.Context, req dto.ProductReq, productId string) (dto.FileData, domain.Product, int, error) {
	file, code, err := ds.fileService.GetFileId(ctx, req.FileId)
	if err != nil {
		return dto.FileData{}, domain.Product{}, code, err
	}
	if file.FileID == "" {
		return dto.FileData{}, domain.Product{}, http.StatusNotFound, errors.New("not found")
	}

	product, err := ds.productRepository.FindById(ctx, productId)
	if err != nil {
		return dto.FileData{}, domain.Product{}, http.StatusInternalServerError, err
	}
	if product.ProductID == 0 {
		return dto.FileData{}, domain.Product{}, http.StatusNotFound, errors.New("not found")
	}

	return file, product, 200, nil
}
