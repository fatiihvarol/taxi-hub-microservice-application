package routes

import "github.com/gofiber/fiber/v2"

func RegisterAllRoutes(app *fiber.App) {
    RegisterAuthRoutes(app)
    RegisterDriverRoutes(app)
    // diğer servisleri eklemek için burada çağır
}
