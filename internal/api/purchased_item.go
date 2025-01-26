package api

import (
	"context"
	"time"

	"github.com/bayuuat/tutuplapak/internal/middleware"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

type purchasedItemApi struct {
	purchasedItemService service.PurchasedItemService
}

func NewPurchasedItem(app *fiber.App,
	purchasedItemService service.PurchasedItemService) {

	da := purchasedItemApi{
		purchasedItemService: purchasedItemService,
	}

	user := app.Group("/v1/purchasedItem")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreatePurchasedItem)
	user.Get("/", da.GetPurchasedItems)
	user.Patch("/:id?", da.UpdatePurchasedItem)
	user.Delete("/:id?", da.DeletePurchasedItem)
}

func (da purchasedItemApi) GetPurchasedItems(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchasedItemApi) CreatePurchasedItem(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchasedItemApi) UpdatePurchasedItem(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}

func (da purchasedItemApi) DeletePurchasedItem(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}
