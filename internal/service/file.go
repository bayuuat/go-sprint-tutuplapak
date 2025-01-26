package service

import (
	"context"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
)

type FileService struct {
	cnf            *config.Config
	fileRepository repository.FileRepository
}

type FileServicer interface {
	GetFilesWithFilter(ctx context.Context, filter dto.FileFilter, userId string) ([]dto.FileData, int, error)
	GetFileId(ctx context.Context, id string) (bool, error)
	CreateFile(ctx context.Context, req dto.FileReq, userId string) (dto.FileData, int, error)
	PatchFile(ctx context.Context, req dto.UpdateFileReq, userId, id string) (dto.FileData, int, error)
	DeleteFile(ctx context.Context, user_id, id string) (dto.FileData, int, error)
}

func NewFile(cnf *config.Config,
	fileRepository repository.FileRepository) FileService {
	return FileService{
		cnf:            cnf,
		fileRepository: fileRepository,
	}
}

func (ds FileService) GetFilesWithFilter(ctx context.Context, filter dto.FileFilter, userId string) ([]dto.FileData, int, error) {
	return []dto.FileData{}, 0, nil
}

func (ds *FileService) CreateFile(ctx context.Context, req dto.FileReq, userId string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) PatchFile(ctx context.Context, req dto.UpdateFileReq, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) DeleteFile(ctx context.Context, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) GetFileId(ctx context.Context, id string) (domain.File, error) {
	file, err := ds.fileRepository.FindById(ctx, "", id)
	if err != nil {
		return domain.File{}, err
	}

	return file, nil
}
