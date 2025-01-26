package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Put(ctx context.Context, product dto.Product) (err error)
	FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter) ([]domain.Product, error)
	FindById(ctx context.Context, id string) (domain.Product, error)
	Delete(ctx context.Context, id string) error
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
	return &domain.Product{}, errors.New("not implemented")
}

func (d productRepository) Put(ctx context.Context, product dto.Product) (err error) {
	file := d.db.Insert("products").Rows(goqu.Record{
		"name":     product.Name,
		"category": product.Category,
		"qty":      product.Qty,
		"price":    product.Price,
		"sku":      product.Sku,
		"file_id":  product.FileId,
	}).Executor()

	_, err = file.ExecContext(ctx)
	return err
}

func (d productRepository) FindById(ctx context.Context, id string) (product domain.Product, err error) {
	dataset := d.db.From("products").Where(goqu.Ex{
		"product_id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &product)
	return product, err
}

func (d productRepository) Delete(ctx context.Context, id string) error {

	return errors.New("not implemented")
}

func (d productRepository) FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter) ([]domain.Product, error) {
	return []domain.Product{}, errors.New("not implemented")
}
