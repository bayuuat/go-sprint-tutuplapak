package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, userId string, product goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter, userId string) ([]domain.Product, error)
	FindById(ctx context.Context, userId, id string) (domain.Product, error)
	Delete(ctx context.Context, userId, id string) error
}

type productRepository struct {
	db *goqu.Database
}

func NewProduct(db *sql.DB) ProductRepository {
	return &productRepository{
		db: goqu.New("default", db),
	}
}

func (d productRepository) Save(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return nil, nil
}

func (d productRepository) Update(ctx context.Context, userId string, product goqu.Record) error {
	return nil
}

func (d productRepository) FindById(ctx context.Context, userId, id string) (product domain.Product, err error) {
	return nil, nil
}

func (d productRepository) Delete(ctx context.Context, userId, id string) error {
	return nil
}

func (d productRepository) FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter, userId string) ([]domain.Product, error) {
	return nil, nil
}
