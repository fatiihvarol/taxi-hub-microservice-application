package controllers

import (
	"auth-service/database"
	"auth-service/dtos"
	"auth-service/services"
	"auth-service/repositories"
	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dtos.RegisterRequest true "User info"
// @Success 200 {object} dtos.RegisterResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	body := new(dtos.RegisterRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(dtos.ErrorResponse{Error: "Invalid body"})
	}

	// Mongo repository
	userRepo := repositories.NewMongoUserRepository(database.UserCollection)

	// Service çağrısı
	resp, errResp := services.RegisterUser(userRepo, body)
	if errResp != nil {
		return c.Status(400).JSON(errResp)
	}

	return c.JSON(resp)
}

