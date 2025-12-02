package controllers

import (
	"driver-service/dtos"
	"driver-service/repositories"
	"driver-service/services"
	"github.com/gofiber/fiber/v2"
	"driver-service/database"
)

type DriverController struct {
	driverService *services.DriverService
}

// Mongo collection burada sağlanmalı
func NewDriverController() *DriverController {
    repo := repositories.NewMongoDriverRepository(database.GetDriverCollection())
    service := services.NewDriverService(repo)
    return &DriverController{
        driverService: service,
    }
}


// Create Driver
// @Summary Create Driver
// @Description Yeni driver oluşturur
// @Tags Drivers
// @Accept json
// @Produce json
// @Param driver body dtos.CreateDriverRequest true "Driver"
// @Success 201 {object} dtos.CreateDriverResponse
// @Router /drivers [post]
func (c *DriverController) CreateDriver(ctx *fiber.Ctx) error {
	var req dtos.CreateDriverRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	resp, err := c.driverService.CreateDriver(&req)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.Status(201).JSON(resp)
}


// Update Driver
// @Summary Update Driver
// @Tags Drivers
// @Accept json
// @Produce json
// @Param id path string true "Driver ID"
// @Param driver body dtos.UpdateDriverRequest true "Driver"
// @Success 200 {object} dtos.UpdateDriverResponse
// @Router /drivers/{id} [put]
func (c *DriverController) UpdateDriver(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var req dtos.UpdateDriverRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	resp, err := c.driverService.UpdateDriver(id, &req)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(resp)
}

// List Drivers
// @Summary List Drivers
// @Tags Drivers
// @Produce json
// @Param page query int false "Page"
// @Param pageSize query int false "Page Size"
// @Success 200 {array} dtos.DriverListItem
// @Router /drivers [get]
func (c *DriverController) GetDrivers(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	size := ctx.QueryInt("pageSize", 10)

	resp, err := c.driverService.ListDrivers(page, size)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(resp)
}

// Nearby Drivers
// @Summary Get Nearby Drivers
// @Tags Drivers
// @Produce json
// @Param lat query float64 true "Latitude"
// @Param lon query float64 true "Longitude"
// @Param taksiType query string true "Taxi type"
// @Success 200 {array} dtos.NearbyDriverResponse
// @Router /drivers/nearby [get]
func (c *DriverController) GetNearbyDrivers(ctx *fiber.Ctx) error {
	lat := ctx.QueryFloat("lat")
	lon := ctx.QueryFloat("lon")
	taxiType := ctx.Query("taksiType")

	resp, err := c.driverService.GetNearby(lat, lon, taxiType)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return ctx.JSON(resp)
}
// HealthCheck
// @Summary Health Check
// @Description Servis durumu kontrolü
// @Tags System
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (c *DriverController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"service": "driver-service",
	})
}
