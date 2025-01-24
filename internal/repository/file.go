package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type FileRepository interface {
	Save(ctx context.Context, file *domain.File) (*domain.File, error)
	Update(ctx context.Context, userId string, file goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.FileFilter, userId string) ([]domain.File, error)
	FindById(ctx context.Context, userId, id string) (domain.File, error)
	Delete(ctx context.Context, userId, id string) error
}

type fileRepository struct {
	db *goqu.Database
}

func NewFile(db *sql.DB) FileRepository {
	return &fileRepository{
		db: goqu.New("default", db),
	}
}

func (d fileRepository) Save(ctx context.Context, file *domain.File) (*domain.File, error) {
	return nil, nil
}

func (d fileRepository) Update(ctx context.Context, userId string, file goqu.Record) error {
	return nil
}

func (d fileRepository) FindById(ctx context.Context, userId, id string) (file domain.File, err error) {
	return nil, nil
}

func (d fileRepository) Delete(ctx context.Context, userId, id string) error {
	return nil
}

func (d fileRepository) FindAllWithFilter(ctx context.Context, filter *dto.FileFilter, userId string) ([]domain.File, error) {
	return nil, nil
}
