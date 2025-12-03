package routes

import (
    "github.com/gofiber/fiber/v2"
    "location-service/controllers"
)

func SetupRoutes(app *fiber.App, c *controller.LocationController) {
    api := app.Group("/location")

    api.Post("/", c.UpdateLocation) 
        app.Get("/ws/drivers", websocket.New(func(conn *websocket.Conn) {
        handler := websocket.DriverSocketHandler(svc)
        handler(conn)
    }))
}
