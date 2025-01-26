package api

import (
	"context"
	"github.com/bayuuat/tutuplapak/dto"
	"time"

	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

type productApi struct {
	productService service.ProductServicer
}

func NewProduct(app *fiber.App,
	productService service.ProductServicer) {

	da := productApi{
		productService: productService,
	}

	user := app.Group("/v1/product")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreateProduct)
	user.Get("/", da.GetProducts)
	user.Put("/:id?", da.UpdateProduct)
	user.Delete("/:id?", da.DeleteProduct)
}

func (da productApi) GetProducts(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) CreateProduct(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) UpdateProduct(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
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
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	_, code, err := da.productService.DeleteProduct(c, "change this", "change this")
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(500).JSON(dto.ErrorResponse{Message: "not implemented"})
}
