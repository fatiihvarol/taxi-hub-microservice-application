package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.ConnectMongo()

	app := fiber.New()

	routes.AuthRoutes(app)

	app.Listen(":" + config.GetEnv("PORT"))
}
