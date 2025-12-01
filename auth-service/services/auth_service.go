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
func LoginUser(repo repositories.UserRepository, req *dtos.LoginRequest) (*dtos.LoginResponse, *dtos.ErrorResponse) {
	user, err := repo.FindByEmail(context.Background(), req.Email)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "User not found"}
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, &dtos.ErrorResponse{Error: "Invalid credentials"}
	}

	accessToken, err := utils.GenerateAccessToken(user.Email)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Could not generate access token"}
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Email)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Could not generate refresh token"}
	}

	return &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func RefreshUserToken(repo repositories.UserRepository, req *dtos.RefreshRequest) (*dtos.LoginResponse, *dtos.ErrorResponse) {
	// Refresh token doğrulama ve yeni access/refresh token üret
	accessToken, refreshToken, err := utils.RefreshJWT(req.RefreshToken)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: err.Error()}
	}

	return &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
