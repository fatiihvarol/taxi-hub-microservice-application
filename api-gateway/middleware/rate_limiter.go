package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiter returns a Fiber middleware for rate limiting
func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,               // 1 dakika içinde maksimum istek
		Expiration: 1 * time.Minute,   // süre
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // IP bazlı limit
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too Many Requests",
			})
		},
	})
}
