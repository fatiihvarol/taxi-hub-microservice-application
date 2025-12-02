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

	// Her endpoint ayrı ayrı role bazlı middleware ile korunabilir
	app.Get("/api/hello", middleware.JWTProtected(authServiceURL)) // token geçerli olan herkes
	app.Get("/api/admin", middleware.JWTProtected(authServiceURL, "admin")) // sadece admin
	app.Get("/api/driver/updateLocation", middleware.JWTProtected(authServiceURL, "driver")) // sadece driver
	app.Get("/api/location/nearby", middleware.JWTProtected(authServiceURL, "customer", "driver")) // customer ve driver

	// Auth service proxy (register/login/validate gibi)
	app.All("/auth/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, authServiceURL+c.OriginalURL())
	})

	log.Fatal(app.Listen(":" + port))
}

