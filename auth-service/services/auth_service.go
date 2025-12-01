package services

import (
	"auth-service/dtos"
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/utils"
	"context"
	"time"
)

func RegisterUser(repo repositories.UserRepository, req *dtos.RegisterRequest) (*dtos.RegisterResponse, *dtos.ErrorResponse) {
	// Hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Password hashing failed"}
	}

	user := models.User{
		Email:     req.Email,
		Password:  hashed,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// DB insert
	err = repo.CreateUser(context.Background(), user)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Database error"}
	}

	return &dtos.RegisterResponse{Message: "registered"}, nil
}
