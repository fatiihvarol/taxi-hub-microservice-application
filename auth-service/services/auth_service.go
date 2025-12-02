package services

import (
	"auth-service/dtos"
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/utils"
	"context"
	"time"
	"auth-service/config"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

func RegisterUser(repo repositories.UserRepository, req *dtos.RegisterRequest) (*dtos.RegisterResponse, *dtos.ErrorResponse) {
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Password hashing failed"}
	}

	role := req.Role
	if role == "" {
		role = "customer" // default role
	}

	user := models.User{
		Email:     req.Email,
		Password:  hashed,
		Role:      role,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

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
		fmt.Println("Fetched user from DB:", user) // <--- Burada Role dolu mu?

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, &dtos.ErrorResponse{Error: "Invalid credentials"}
	}

	// JWT içine role ekle
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Could not generate access token"}
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return nil, &dtos.ErrorResponse{Error: "Could not generate refresh token"}
	}

	return &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role:         user.Role, // opsiyonel: response'a da ekleyebilirsin
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
func ValidateUserToken(token string) (*models.User, error) {
	parsedToken, err := utils.ParseJWT(token, config.GetEnv("JWT_SECRET"))
	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
userID := claims["userId"].(string)
role := claims["role"].(string)


	return &models.User{ID: userID, Role: role}, nil
}
