package controller

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "location-service/dtos"
    "location-service/models"
    "location-service/services"
)

type LocationController struct {
    service *service.LocationService
}

func NewLocationController(s *service.LocationService) *LocationController {
    return &LocationController{s}
}

func (c *LocationController) UpdateLocation(ctx *fiber.Ctx) error {
    var req dtos.UpdateLocationRequest
    if err := ctx.BodyParser(&req); err != nil {
        return fiber.ErrBadRequest
    }

    loc := &models.DriverLocation{
        Lat:       req.Lat,
        Lon:       req.Lon,
        TaxiType:  req.TaxiType,
        UpdatedAt: time.Now().Unix(),
    }

    // driverID artık body’den geliyor
    err := c.service.Update(req.DriverID, loc)
    if err != nil {
        return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(fiber.Map{"status": "ok"})
}


func (c *LocationController) GetNearby(ctx *fiber.Ctx) error {
    lat := ctx.QueryFloat("lat")
    lon := ctx.QueryFloat("lon")
    taxiType := ctx.Query("taksiType")

    res, err := c.service.Nearby(lat, lon, taxiType)
    if err != nil {
        return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(res)
}
