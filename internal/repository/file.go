package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/doug-martin/goqu/v9"
)

type FileRepository interface {
	Save(ctx context.Context, file *domain.File) (*domain.File, error)
	Update(ctx context.Context, userId string, file goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.FileFilter, userId string) ([]domain.File, error)
	FindById(ctx context.Context, id string) (domain.File, error)
	FindByIds(ctx context.Context, ids []string) (files []domain.File, err error)
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
	return &domain.File{}, errors.New("not implemented")
}

func (d fileRepository) Update(ctx context.Context, userId string, file goqu.Record) error {
	return errors.New("not implemented")
}

func (d fileRepository) FindById(ctx context.Context, id string) (file domain.File, err error) {
	dataset := d.db.From("files").Where(goqu.Ex{
		"file_id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &file)
	return file, err
}

func (d fileRepository) FindByIds(ctx context.Context, ids []string) (files []domain.File, err error) {
	err = d.db.From("files").Where(goqu.C("file_id").In(ids)).ScanStructsContext(ctx, &files)
	return files, err
}

func (d fileRepository) Delete(ctx context.Context, userId, id string) error {
	return errors.New("not implemented")
}

func (d fileRepository) FindAllWithFilter(ctx context.Context, filter *dto.FileFilter, userId string) ([]domain.File, error) {
	return []domain.File{}, errors.New("not implemented")
}
