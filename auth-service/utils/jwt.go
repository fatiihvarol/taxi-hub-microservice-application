package utils

import (
	"time"
	"errors"

	"auth-service/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId string, role string) (string, error) {
	secret := config.GetEnv("JWT_SECRET")
	expMinutes := time.Duration(15) * time.Minute

	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(expMinutes).Unix(),

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userId string, role string) (string, error) {
	secret := config.GetEnv("REFRESH_SECRET")
	expHours := time.Duration(720) * time.Hour

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(expHours).Unix(),
		"role":   role,
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

	// role varsa al, yoksa default ata
	role, ok := claims["role"].(string)
	if !ok {
		role = "customer"
	}

	// role ile access token oluştur
	newAccess, err := GenerateAccessToken(userId, role)
	if err != nil {
		return "", "", err
	}

	newRefresh, err := GenerateRefreshToken(userId, role)
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
func ParseJWT(tokenStr string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
}
