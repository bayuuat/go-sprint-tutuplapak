package api

import (
	"context"
	"time"

	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

type purchaseApi struct {
	purchaseService service.PurchaseService
}

func NewPurchase(app *fiber.App,
	purchaseService service.PurchaseService) {

	da := purchaseApi{
		purchaseService: purchaseService,
	}

	user := app.Group("/v1/purchase")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreatePurchase)
	user.Get("/", da.GetPurchases)
	user.Patch("/:id?", da.UpdatePurchase)
	user.Delete("/:id?", da.DeletePurchase)
}

func (da purchaseApi) GetPurchases(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchaseApi) CreatePurchase(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchaseApi) UpdatePurchase(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchaseApi) DeletePurchase(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}
