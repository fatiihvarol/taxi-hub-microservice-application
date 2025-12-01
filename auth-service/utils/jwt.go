package utils

import (
	"time"

	"auth-service/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId string) (string, error) {
	secret := config.GetEnv("JWT_SECRET")
	expMinutes := time.Duration(15) * time.Minute

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(expMinutes).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userId string) (string, error) {
	secret := config.GetEnv("REFRESH_SECRET")
	expHours := time.Duration(720) * time.Hour

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(expHours).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
