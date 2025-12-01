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

// Login godoc
// @Summary Login a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dtos.LoginRequest true "Login info"
// @Success 200 {object} dtos.LoginResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	body := new(dtos.LoginRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(dtos.ErrorResponse{Error: "Invalid body"})
	}

	userRepo := repositories.NewMongoUserRepository(database.UserCollection)
	resp, errResp := services.LoginUser(userRepo, body)
	if errResp != nil {
		return c.Status(400).JSON(errResp)
	}

	return c.JSON(resp)
}
// RefreshToken godoc
// @Summary Refresh JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dtos.RefreshRequest true "Refresh token info"
// @Success 200 {object} dtos.LoginResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Router /auth/refresh [post]
func RefreshToken(c *fiber.Ctx) error {
	body := new(dtos.RefreshRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(dtos.ErrorResponse{Error: "Invalid body"})
	}

	// Mongo repository
	userRepo := repositories.NewMongoUserRepository(database.UserCollection)

	// Service çağrısı
	resp, errResp := services.RefreshUserToken(userRepo, body)
	if errResp != nil {
		return c.Status(400).JSON(errResp)
	}

	return c.JSON(resp)
}





