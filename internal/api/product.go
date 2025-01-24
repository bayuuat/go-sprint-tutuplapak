package api

import (
	"context"
	"time"

	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

type productApi struct {
	productService service.ProductService
}

func NewProduct(app *fiber.App,
	productService service.ProductService) {

	da := productApi{
		productService: productService,
	}

	user := app.Group("/v1/product")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreateProduct)
	user.Get("/", da.GetProducts)
	user.Patch("/:id?", da.UpdateProduct)
	user.Delete("/:id?", da.DeleteProduct)
}

func (da productApi) GetProducts(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) CreateProduct(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) UpdateProduct(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da productApi) DeleteProduct(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}
