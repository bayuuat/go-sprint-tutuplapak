package main

import (
	"github.com/bayuuat/tutuplapak/internal/api"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/connection"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cnf := config.Get()
	app := fiber.New()

	dbConnection := goqu.New("postgres", connection.GetDatabase(cnf.Database))
	dbConnection.Logger(log.Default())

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewUser(cnf, userRepository)
	api.NewUser(app, authService)

	fileRepository := repository.NewFile(dbConnection)
	fileService := service.NewFile(cnf, fileRepository)
	// api.NewFile()???

	productRepository := repository.NewProduct(dbConnection)
	productService := service.NewProductServicer(cnf, productRepository, fileService)
	api.NewProduct(app, productService)

	purchasedItemRepository := repository.NewPurchasedItem(dbConnection)
	purchasedItemService := service.NewPurchasedItemServicer(cnf, purchasedItemRepository, productService, fileService)

	purchaseRepository := repository.NewPurchase(dbConnection)
	purchaseService := service.NewPurchaseServicer(cnf, purchaseRepository, purchasedItemService, authService)
	api.NewPurchase(app, purchaseService)

	api.NewAws(app)

	// TODO TOMORROW: Add stock check, omitempty db, senderValidation
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
