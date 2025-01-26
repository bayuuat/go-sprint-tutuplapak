package main

import (
	"github.com/bayuuat/tutuplapak/internal/api"
	"github.com/bayuuat/tutuplapak/internal/config"
	"github.com/bayuuat/tutuplapak/internal/connection"
	"github.com/bayuuat/tutuplapak/internal/repository"
	"github.com/bayuuat/tutuplapak/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	app := fiber.New()
	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewUser(cnf, userRepository)
	api.NewUser(app, authService)

	fileRepository := repository.NewFile(dbConnection)
	fileService := service.NewFile(cnf, fileRepository)

	productRepository := repository.NewProduct(dbConnection)
	productService := service.NewProductServicer(cnf, productRepository, fileService)
	api.NewProduct(app, productService)

	api.NewAws(app)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
