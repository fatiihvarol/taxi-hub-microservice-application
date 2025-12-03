package controllers

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "location-service/dtos"
    "location-service/models"
    "location-service/services" // services olarak import
)

type LocationController struct {
    service *services.LocationService // services olarak kullan
}

func NewLocationController(s *services.LocationService) *LocationController {
    return &LocationController{s}
}
// UpdateLocation godoc
// @Summary Update driver location
// @Description Update the location of a driver
// @Tags location
// @Accept json
// @Produce json
// @Param location body dtos.UpdateLocationRequest true "Location payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /location [post]
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

    err := c.service.Update(req.DriverID, loc)
    if err != nil {
        return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(fiber.Map{"status": "ok"})
}
// GetNearby godoc
// @Summary Get nearby drivers
// @Description Get drivers near the specified location
// @Tags location
// @Accept json
// @Produce json
// @Param lat query number true "Latitude"
// @Param lon query number true "Longitude"
// @Param taksiType query string true "Taxi type"
// @Success 200 {array} models.DriverLocation
// @Failure 500 {object} map[string]string
// @Router /location/nearby [get]
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
