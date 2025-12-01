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
	// .env yükle
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set")
	}

	app := fiber.New()
	app.Use(logger.New())

	// JWT Middleware - tüm /api/* endpointleri korur
	app.Use("/api", middleware.JWTProtected(jwtSecret))

	// Örnek protected endpoint
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, authorized user!"})
	})

	// Auth-service route’larını proxy ile yönlendir
	app.All("/auth/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, authServiceURL+c.OriginalURL())
	})

	log.Fatal(app.Listen(":" + port))
}
