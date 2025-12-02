package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
    "io"
	"github.com/gofiber/fiber/v2"
)

// JWTProtected returns a Fiber middleware that validates the JWT token
// and optionally checks allowed roles
func JWTProtected(authServiceURL string, allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Println("[JWTProtected] Missing Authorization header")
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token format"})
		}

		// token değişkenini kullanılabilir yap
		_ = parts[1]

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
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[JWTProtected] Failed to read auth service response: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}

		var tokenData map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &tokenData); err != nil {
			log.Printf("[JWTProtected] Failed to parse auth service response: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}

		userRole, ok := tokenData["role"].(string)
		if !ok {
			userRole = ""
		}

		if len(allowedRoles) > 0 {
			authorized := false
			for _, role := range allowedRoles {
				if role == userRole {
					authorized = true
					break
				}
			}
			if !authorized {
				return c.Status(403).JSON(fiber.Map{"error": "Forbidden: insufficient role"})
			}
		}

		return c.Next()
	}
}

