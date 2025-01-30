package api

import (
	"context"
	"github.com/bayuuat/tutuplapak/dto"
	"time"

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

	user.Post("/", da.CreatePurchase)
	user.Post("/:id?", da.CreatePayment)
}


func (da purchaseApi) CreatePayment(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	return ctx.Status(200).JSON(nil)
}
  

func (da purchaseApi) CreatePurchase(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var purchase dto.PurchaseReq
	if err := ctx.BodyParser(&purchase); err != nil {
		return ctx.Status(400).JSON(fiber.Map{})
	}

	res, code, err := da.purchaseService.CreatePurchase(c, purchase)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
