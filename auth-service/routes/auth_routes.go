package routes

import (
    "auth-service/controllers"
    "github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/auth/register", controllers.Register)

}
