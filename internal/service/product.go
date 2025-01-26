package service

import (
	"context"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type productService struct {
	cnf               *config.Config
	productRepository repository.ProductRepository
}

type ProductService interface {
	GetProductsWithFilter(ctx context.Context, filter dto.ProductFilter, userId string) ([]dto.ProductData, int, error)
	CreateProduct(ctx context.Context, req dto.ProductReq, userId string) (dto.ProductData, int, error)
	PatchProduct(ctx context.Context, req dto.UpdateProductReq, userId, id string) (dto.ProductData, int, error)
	DeleteProduct(ctx context.Context, user_id, id string) (dto.ProductData, int, error)
}

func NewProduct(cnf *config.Config,
	productRepository repository.ProductRepository) ProductService {
	return &productService{
		cnf:               cnf,
		productRepository: productRepository,
	}
}

func (ds productService) GetProductsWithFilter(ctx context.Context, filter dto.ProductFilter, userId string) ([]dto.ProductData, int, error) {
	return []dto.ProductData{}, 200, nil
}

func (ds *productService) CreateProduct(ctx context.Context, req dto.ProductReq, userId string) (dto.ProductData, int, error) {
	return dto.ProductData{}, 200, nil
}

func (ds productService) PatchProduct(ctx context.Context, req dto.UpdateProductReq, userId, id string) (dto.ProductData, int, error) {
	return dto.ProductData{}, 200, nil
}

func (ds productService) DeleteProduct(ctx context.Context, userId, id string) (dto.ProductData, int, error) {
	return dto.ProductData{}, 200, nil
}
