package routes

import (
    "github.com/gofiber/fiber/v2"
    "location-service/controller"
)

func SetupRoutes(app *fiber.App, c *controller.LocationController) {
    api := app.Group("/location")

    api.Post("/:driverId", c.UpdateLocation)
    api.Get("/nearby", c.GetNearby)
}
