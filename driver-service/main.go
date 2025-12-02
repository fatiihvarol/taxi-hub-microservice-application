package main

import (
	"driver-service/config"
	"driver-service/controllers"
	"driver-service/database"
	"driver-service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	config.LoadEnv()
	database.ConnectMongo()

	// Fiber app
	app := fiber.New()
	app.Get("/docs/*", swagger.HandlerDefault)

	// Controller initialize
	driverController := controllers.NewDriverController()

	// Driver routes
	routes.DriverRoutes(app, driverController)

	// Start server
	app.Listen(":" + config.GetEnv("PORT"))
}
