package controllers

import (
    "auth-service/database"
    "auth-service/dtos"
    "auth-service/models"
    "auth-service/utils"
    "context"
    "time"

    "github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dtos.RegisterRequest true "User info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dtos.ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
    body := new(dtos.RegisterRequest)
    if err := c.BodyParser(body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
    }

    // Hash password
    hashed, _ := utils.HashPassword(body.Password)

    user := models.User{
        Email:     body.Email,
        Password:  hashed,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := database.UserCollection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "DB error"})
    }

    return c.JSON(fiber.Map{"message": "registered"})
}
