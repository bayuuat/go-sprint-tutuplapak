package service

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/bayuuat/tutuplapak/domain"
	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"github.com/bayuuat/tutuplapak/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterEmail(ctx context.Context, req dto.AuthEmailReq) (dto.AuthResponse, int, error)
	LoginEmail(ctx context.Context, req dto.AuthEmailReq) (dto.AuthResponse, int, error)
	RegisterPhone(ctx context.Context, req dto.AuthPhoneReq) (dto.AuthResponse, int, error)
	LoginPhone(ctx context.Context, req dto.AuthPhoneReq) (dto.AuthResponse, int, error)
	GetUser(ctx context.Context, email string) (dto.UserData, int, error)
	PatchUser(ctx context.Context, req dto.UpdateUser, email string) (dto.UserData, int, error)
}

type userService struct {
	cnf            *config.Config
	userRepository repository.UserRepository
}

func NewUser(cnf *config.Config,
	userRepository repository.UserRepository) UserService {
	return &userService{
		cnf:            cnf,
		userRepository: userRepository,
	}
}

func (a userService) RegisterEmail(ctx context.Context, req dto.AuthEmailReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id != "" {
		return dto.AuthResponse{}, http.StatusConflict, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	newUser := domain.User{
		Id:       uuid.New().String(),
		Email:    &req.Email,
		Password: string(hashedPassword),
	}

	err = a.userRepository.Save(ctx, &newUser)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	user = newUser

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: *user.Email,
		Phone: nil,
		Token: token,
	}, http.StatusCreated, nil
}

func (a userService) RegisterPhone(ctx context.Context, req dto.AuthPhoneReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByPhone(ctx, req.Phone)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id != "" {
		return dto.AuthResponse{}, http.StatusConflict, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	newUser := domain.User{
		Id:       uuid.New().String(),
		Phone:    &req.Phone,
		Password: string(hashedPassword),
	}

	err = a.userRepository.Save(ctx, &newUser)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	user = newUser

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: *user.Email,
		Phone: *user.Phone,
		Token: token,
	}, http.StatusCreated, nil
}

func (a userService) LoginEmail(ctx context.Context, req dto.AuthEmailReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, http.StatusNotFound, domain.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.AuthResponse{}, http.StatusUnauthorized, domain.ErrInvalidCredential
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: *user.Email,
		Token: token,
	}, http.StatusOK, nil
}

func (a userService) LoginPhone(ctx context.Context, req dto.AuthPhoneReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Phone)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, http.StatusNotFound, domain.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.AuthResponse{}, http.StatusUnauthorized, domain.ErrInvalidCredential
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: *user.Email,
		Token: token,
	}, http.StatusOK, nil
}

func (a userService) GetUser(ctx context.Context, email string) (dto.UserData, int, error) {
	_, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	return dto.UserData{}, http.StatusOK, nil
}

func (a userService) PatchUser(ctx context.Context, req dto.UpdateUser, id string) (dto.UserData, int, error) {
	user, err := a.userRepository.FindById(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	err = a.userRepository.Update(ctx, &user)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserData{}, http.StatusInternalServerError, err
	}

	return dto.UserData{}, 200, nil
}
