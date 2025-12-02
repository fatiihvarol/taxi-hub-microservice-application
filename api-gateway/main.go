package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"api-gateway/elastic"
	"api-gateway/middleware"
	"api-gateway/routes"

)

func main() {
	// .env yükle
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Elasticsearch initialize
	elastic.InitElastic()

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Fiber
	app := fiber.New()
	app.Use(middleware.APIKeyMiddleware()) 

	app.Use(middleware.RateLimiter())

	// Normal log (console log)
	app.Use(logger.New())

	// Global Elasticsearch logger
	app.Use(middleware.GlobalLogger)

    routes.RegisterAllRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
