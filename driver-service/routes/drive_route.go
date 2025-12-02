package routes

import (
	"driver-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func DriverRoutes(app *fiber.App, controller *controllers.DriverController) {
	driver := app.Group("/drivers")

	driver.Post("/", controller.CreateDriver)
	driver.Put("/:id", controller.UpdateDriver)
	driver.Get("/", controller.GetDrivers)
	driver.Get("/nearby", controller.GetNearbyDrivers)
	driver.Get("/health", controller.HealthCheck)
}
