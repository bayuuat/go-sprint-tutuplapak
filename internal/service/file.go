package service

import (
	"context"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type fileService struct {
	cnf            *config.Config
	fileRepository repository.FileRepository
}

type FileService interface {
	GetFilesWithFilter(ctx context.Context, filter dto.FileFilter, userId string) ([]dto.FileData, int, error)
	CreateFile(ctx context.Context, req dto.FileReq, userId string) (dto.FileData, int, error)
	PatchFile(ctx context.Context, req dto.UpdateFileReq, userId, id string) (dto.FileData, int, error)
	DeleteFile(ctx context.Context, user_id, id string) (dto.FileData, int, error)
}

func NewFile(cnf *config.Config,
	fileRepository repository.FileRepository) FileService {
	return &fileService{
		cnf:            cnf,
		fileRepository: fileRepository,
	}
}

func (ds fileService) GetFilesWithFilter(ctx context.Context, filter dto.FileFilter, userId string) ([]dto.FileData, int, error) {
	return []dto.FileData{}, 200, nil
}

func (ds *fileService) CreateFile(ctx context.Context, req dto.FileReq, userId string) (dto.FileData, int, error) {
	return dto.FileData{}, 200, nil
}

func (ds fileService) PatchFile(ctx context.Context, req dto.UpdateFileReq, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 200, nil
}

func (ds fileService) DeleteFile(ctx context.Context, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 200, nil
}
