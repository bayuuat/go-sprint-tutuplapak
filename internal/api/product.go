package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bayuuat/tutuplapak/dto"
	"github.com/golang-jwt/jwt/v5"

	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

type productApi struct {
	productService service.ProductServicer
}

func NewProduct(app *fiber.App, productService service.ProductServicer) {

	da := productApi{
		productService: productService,
	}
	// Rute ini tidak membutuhkan login, jadi kita batalkan middleware JWT untuk rute GET ini
	// userNoAuth := app.Group("/v1/products")
	// userNoAuth.Get("/", da.GetProducts) // Rute ini akan tetap bisa diakses tanpa login

	user := app.Group("/v1/product")

	user.Get("/", da.GetProducts)

	// Terapkan middleware JWT hanya untuk rute yang membutuhkan login
	user.Use(middleware.JWTProtected) // Ini akan diterapkan pada semua rute dalam grup ini

	// Rute ini membutuhkan login
	user.Post("/", da.CreateProduct)
	user.Put("/:id?", da.UpdateProduct)
	user.Delete("/:id?", da.DeleteProduct)

}

func (da productApi) GetProducts(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var filter dto.ProductFilter
	if err := ctx.QueryParser(&filter); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	res, code, err := da.productService.GetProductsWithFilter(ctx.Context(), filter, "change this")

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da productApi) CreateProduct(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) UpdateProduct(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	q := ctx.Query("id")
	if q == "" {
		return ctx.Status(404).JSON(fiber.Map{})
	}

	var product dto.ProductReq
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	res, code, err := da.productService.PutProduct(c, product, q)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(200).JSON(res)
}

func (da productApi) DeleteProduct(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	id := ctx.Params("id")

	res, code, err := da.productService.DeleteProduct(c, userId, id)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
