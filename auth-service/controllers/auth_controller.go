package controllers

import (
	"context"
	"time"

	"auth-service/database"
	"auth-service/models"
	"auth-service/utils"
	"github.com/gofiber/fiber/v2"
)

type RegisterBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	body := new(RegisterBody)
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
