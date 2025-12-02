package middleware

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

func APIKeyMiddleware() fiber.Handler {
    apiKey := os.Getenv("API_KEY")
    return func(c *fiber.Ctx) error {
        reqKey := c.Get("X-API-Key")
        if reqKey != apiKey {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid API Key",
            })
        }
        return c.Next()
    }
}
