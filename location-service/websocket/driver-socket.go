package websocket

import (
    "github.com/gofiber/websocket/v2"
)

func DriverSocketHandler() func(*websocket.Conn) {
    return func(c *websocket.Conn) {
        for {
            _, msg, err := c.ReadMessage()
            if err != nil {
                return
            }

            // TODO: gelen konumu parse et service.Update() çağır
        }
    }
}
