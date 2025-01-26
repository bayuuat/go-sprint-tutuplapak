package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/bayuuat/tutuplapak/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type authApi struct {
	authService service.UserService
}

func NewUser(app *fiber.App, authService service.UserService) {

	ha := authApi{
		authService: authService,
	}

	app.Post("/v1/register/email", ha.RegisterEmail)
	app.Post("/v1/login/email", ha.LoginEmail)
	app.Post("/v1/register/phone", ha.RegisterPhone)
	app.Post("/v1/login/phone", ha.LoginPhone)

	user := app.Group("/v1/user")

	user.Use(middleware.JWTProtected)
	user.Get("/", ha.GetUser)
	user.Patch("/", ha.UpdateUser)
}

func (a authApi) RegisterEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthEmailReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := a.authService.RegisterEmail(c, req)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (a authApi) LoginEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthEmailReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := a.authService.LoginEmail(c, req)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (a authApi) RegisterPhone(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthPhoneReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := a.authService.RegisterPhone(c, req)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (a authApi) LoginPhone(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthPhoneReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := a.authService.LoginPhone(c, req)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (a authApi) GetUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	// Get email claims
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	res, code, err := a.authService.GetUser(c, email)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (a authApi) UpdateUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var req dto.UpdateUser
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := a.authService.PatchUser(c, req, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
