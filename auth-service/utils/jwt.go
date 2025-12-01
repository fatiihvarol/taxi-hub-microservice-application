package utils

import (
	"time"
	"errors"

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
func RefreshJWT(refreshToken string) (string, string, error) {
	secret := config.GetEnv("REFRESH_SECRET")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", "", errors.New("invalid user id in token")
	}

	newAccess, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	newRefresh, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return newAccess, newRefresh, nil
}

func ValidateJWT(tokenStr string, secret string) (bool, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}
