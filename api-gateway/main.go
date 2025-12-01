package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/joho/godotenv"
	"api-gateway/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	app := fiber.New()
	app.Use(logger.New())

	// JWT korumalı endpoint
	app.Use("/api", middleware.JWTProtected(authServiceURL))

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, authorized user!"})
	})

	// Auth service proxy
	app.All("/auth/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, authServiceURL+c.OriginalURL())
	})

	log.Fatal(app.Listen(":" + port))
}
