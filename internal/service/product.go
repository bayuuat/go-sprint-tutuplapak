package service

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type ProductService struct {
	cnf               *config.Config
	productRepository repository.ProductRepository
	fileService       FileService
}

type ProductServicer interface {
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

func (ds ProductService) GetProductsWithFilter(ctx context.Context, filter dto.ProductFilter, userId string) ([]dto.ProductData, int, error) {
	products, err := ds.productRepository.FindAllWithFilter(ctx, &filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var productData []dto.ProductData
	for _, product := range products {
		file, err := ds.fileService.fileRepository.GetFile(ctx, product.FileID)
		if err != nil {
			return nil, http.StatusNotFound, err
		}

		productData = append(productData, dto.ProductData{
			ProductId:        strconv.Itoa(product.ProductID),
			Name:             product.Name,
			Category:         product.Category,
			Qty:              product.Qty,
			Price:            int(product.Price),
			Sku:              product.SKU,
			FileId:           strconv.Itoa(file.FileID),
			FileUri:          file.FileURI,
			FileThumbnailUri: file.FileThumbnailURI,
			CreatedAt:        product.CreatedAt.Format(time.RFC3339Nano),
			UpdatedAt:        product.UpdatedAt.Format(time.RFC3339Nano),
		})
	}

	return productData, http.StatusOK, nil

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
		FileId:           strconv.Itoa(file.FileID),
		FileUri:          file.FileURI,
		FileThumbnailUri: file.FileThumbnailURI,
		CreatedAt:        time.Now().Format(time.RFC3339Nano),
		UpdatedAt:        time.Now().Format(time.RFC3339Nano),
	}, 200, nil
}

func (ds ProductService) validateProduct(ctx context.Context, req dto.ProductReq, productId string) (domain.File, domain.Product, int, error) {
	file, err := ds.fileService.GetFileId(ctx, req.FileId)
	if err != nil {
		return domain.File{}, domain.Product{}, http.StatusInternalServerError, err
	}
	if file.FileID == 0 {
		return domain.File{}, domain.Product{}, http.StatusNotFound, errors.New("not found")
	}

	product, err := ds.productRepository.FindById(ctx, productId)
	if err != nil {
		return domain.File{}, domain.Product{}, http.StatusInternalServerError, err
	}
	if product.ProductID == 0 {
		return domain.File{}, domain.Product{}, http.StatusNotFound, errors.New("not found")
	}

	return file, product, 200, nil
}

func (ds ProductService) DeleteProduct(ctx context.Context, userId, id string) (dto.ProductData, int, error) {
	if id == "" {
		return dto.ProductData{}, http.StatusNotFound, errors.New("Not found product ID")
	}

	product, err := ds.productRepository.FindById(ctx, id)

	if err != nil {
		return dto.ProductData{}, http.StatusInternalServerError, err
	}

	if product.ProductID == 0 {
		return dto.ProductData{}, http.StatusInternalServerError, domain.ErrActivityNotFound
	}

	err = ds.productRepository.Delete(ctx, userId, id)
	if err != nil {
		return dto.ProductData{}, http.StatusInternalServerError, err
	}

	return dto.ProductData{}, http.StatusOK, err
}
