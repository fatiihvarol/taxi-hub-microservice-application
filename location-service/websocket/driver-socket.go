package websocket

import (
    "encoding/json"
    "time"

    "github.com/gofiber/websocket/v2"
    "location-service/dtos"
    "location-service/models"
    "location-service/services" 


func DriverSocketHandler(service *service.LocationService) func(*websocket.Conn) {
    return func(c *websocket.Conn) {
        defer c.Close()
        for {
            _, msg, err := c.ReadMessage()
            if err != nil {
                return
            }

            var req dtos.UpdateLocationRequest
            if err := json.Unmarshal(msg, &req); err != nil {
                continue
            }

            loc := &models.DriverLocation{
                Lat:       req.Lat,
                Lon:       req.Lon,
                TaxiType:  req.TaxiType,
                UpdatedAt: time.Now().Unix(),
            }

            _ = service.Update(req.DriverID, loc) // hata loglanabilir
        }
    }
}

