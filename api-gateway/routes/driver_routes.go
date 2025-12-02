package routes

import (
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/proxy"
    "api-gateway/middleware"
)

func RegisterDriverRoutes(app *fiber.App) {
    driverServiceURL := os.Getenv("DRIVER_SERVICE_URL")
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    base := "/drivers"

    // tüm driver endpointleri JWT ile korunacak
    app.All(base+"/*",
        middleware.JWTProtected(authServiceURL, "admin", "driver"),
        func(c *fiber.Ctx) error {
            return proxy.Do(c, driverServiceURL+c.OriginalURL())
        },
    )

    // health check ayrı route, sadece admin
    app.Get(base+"/health",
        middleware.JWTProtected(authServiceURL, "admin"),
        func(c *fiber.Ctx) error {
            return proxy.Do(c, driverServiceURL+c.OriginalURL())
        },
    )
}
