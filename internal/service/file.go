package service

import (
	"context"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"net/http"
	"strconv"
)

type FileService struct {
	cnf            *config.Config
	fileRepository repository.FileRepository
}

type FileServicer interface {
	GetFilesWithFilter(ctx context.Context, filter dto.FileFilter, userId string) ([]dto.FileData, int, error)
	GetFileId(ctx context.Context, id string) (dto.FileData, int, error)
	GetFileIds(ctx context.Context, ids []string) ([]dto.FileData, int, error)
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

func (ds FileService) CreateFile(ctx context.Context, req dto.FileReq, userId string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) PatchFile(ctx context.Context, req dto.UpdateFileReq, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) DeleteFile(ctx context.Context, userId, id string) (dto.FileData, int, error) {
	return dto.FileData{}, 0, nil
}

func (ds FileService) GetFileId(ctx context.Context, id string) (dto.FileData, int, error) {
	file, err := ds.fileRepository.FindById(ctx, id)
	if err != nil {
		return dto.FileData{}, http.StatusInternalServerError, err
	}

	return dto.FileData{
		FileID:           strconv.Itoa(file.FileID),
		FileURI:          file.FileURI,
		FileThumbnailURI: file.FileThumbnailURI,
	}, http.StatusOK, nil
}

func (ds FileService) GetFileIds(ctx context.Context, ids []string) ([]dto.FileData, int, error) {
	files, err := ds.fileRepository.FindByIds(ctx, ids)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	filesData := make([]dto.FileData, len(ids))

	for i, file := range files {
		filesData[i] = dto.FileData{
			FileID:           strconv.Itoa(file.FileID),
			FileURI:          file.FileURI,
			FileThumbnailURI: file.FileThumbnailURI,
		}
	}

	return filesData, http.StatusOK, nil
}
