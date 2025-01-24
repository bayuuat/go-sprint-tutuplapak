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

	departmentRepository := repository.NewActivity(dbConnection)
	activityTypesRepository := repository.NewActivityType(dbConnection)
	departmentService := service.NewActivity(cnf, departmentRepository, activityTypesRepository)
	api.NewActivity(app, departmentService)

	api.NewAws(app)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
