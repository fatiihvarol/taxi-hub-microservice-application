package routes

import (
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/proxy"
)

func RegisterAuthRoutes(app *fiber.App) {
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")

    app.All("/auth/*", func(c *fiber.Ctx) error {
        return proxy.Do(c, authServiceURL+c.OriginalURL())
    })
}
