package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"api-gateway/utils"
)

// JWTProtected middleware fonksiyonu, JWT doğrular
func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization") // Bearer <token>
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// "Bearer <token>" kısmını ayır
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		token := tokenParts[1]

		valid, err := utils.ValidateJWT(token, secret)
		if err != nil || !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		return c.Next()
	}
}
