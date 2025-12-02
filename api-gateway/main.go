package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/joho/godotenv"

	"api-gateway/elastic"
	"api-gateway/middleware"
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

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	// Fiber
	app := fiber.New()

	// Normal log (console log)
	app.Use(logger.New())

	// Global Elasticsearch logger
	app.Use(middleware.GlobalLogger())

	// Protected routes
	app.Get("/api/hello", middleware.JWTProtected(authServiceURL))
	app.Get("/api/admin", middleware.JWTProtected(authServiceURL, "admin"))
	app.Get("/api/driver/updateLocation", middleware.JWTProtected(authServiceURL, "driver"))
	app.Get("/api/location/nearby", middleware.JWTProtected(authServiceURL, "customer", "driver"))

	// Auth service proxy
	app.All("/auth/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, authServiceURL+c.OriginalURL())
	})

	log.Fatal(app.Listen(":" + port))
}
