package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected(authServiceURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("[JWTProtected] Missing Authorization header")
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		// Auth-service’e doğrulama isteği at
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/validate", authServiceURL), nil)
		if err != nil {
			log.Printf("[JWTProtected] Failed to create request: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[JWTProtected] Auth service unreachable: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Auth service unreachable"})
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Printf("[JWTProtected] Invalid token. Auth service responded with status: %d\n", resp.StatusCode)
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		log.Println("[JWTProtected] Token validated successfully")
		return c.Next()
	}
}
