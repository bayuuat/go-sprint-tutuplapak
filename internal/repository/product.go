package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type ProductRepository interface {
	Save(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Put(ctx context.Context, product dto.Product) (err error)
	FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter) ([]domain.Product, error)
	FindById(ctx context.Context, id string) (domain.Product, error)
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

func (d productRepository) Delete(ctx context.Context, userId, id string) error {
	query := d.db.From("products").Where(goqu.Ex{
		"product_id": id,
	})

	sql, _, err := query.Delete().ToSQL()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return fmt.Errorf("Error generating SQL: %w", err)
	}

	_, err = d.db.Exec(sql)
	if err != nil {
		log.Println("Error executing SQL:", err)
		return fmt.Errorf("Error executing SQL: %w", err)
	}

	return err
}

func (d productRepository) FindAllWithFilter(ctx context.Context, filter *dto.ProductFilter) ([]domain.Product, error) {
	query := d.db.From("products")

	if filter.Limit > 0 {
		query = query.Limit(uint(filter.Limit))
	} else {
		query = query.Limit(5)
	}
	if filter.Offset > 0 {
		query = query.Offset(uint(filter.Offset))
	} else {
		query = query.Offset(0)
	}

	if filter.ProductID != "" {
		query = query.Where(goqu.Ex{"product_id": filter.ProductID})
	}

	if filter.SKU != "" {
		query = query.Where(goqu.Ex{"sku": filter.SKU})
	}

	if filter.Category != "" {
		query = query.Where(goqu.Ex{"category": filter.Category})
	}

	if filter.SortBy != "" {
		switch {
		case filter.SortBy == "newest":
			query = query.Order(goqu.C("updated_at").Desc(), goqu.C("created_at").Desc())

		case filter.SortBy == "cheapest":
			query = query.Order(goqu.C("price").Asc())
		}
	}

	var products []domain.Product
	err := query.ScanStructsContext(ctx, &products)
	return products, err
}
