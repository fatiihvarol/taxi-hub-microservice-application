package routes

import (
    "github.com/gofiber/fiber/v2"
    ws "github.com/gofiber/websocket/v2"

    "location-service/websocket"
    "location-service/services"
    "location-service/controllers"
)

func SetupRoutes(app *fiber.App, c *controllers.LocationController, svc *services.LocationService) {
    api := app.Group("/location")
    api.Post("/", c.UpdateLocation)
    api.Get("/nearby", c.GetNearby)

    // WebSocket route
    app.Get("/ws/drivers", ws.New(func(conn *ws.Conn) {
        handler := websocket.DriverSocketHandler(svc)
        handler(conn)
    }))
}
