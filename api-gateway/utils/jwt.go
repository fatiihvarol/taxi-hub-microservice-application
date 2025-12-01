package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateJWT token geçerliliğini kontrol eder
func ValidateJWT(tokenString string, secret string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Doğru imzalama yöntemi
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Opsiyonel: token süresini kontrol et
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return false, errors.New("token expired")
			}
		}
		return true, nil
	}

	return false, errors.New("invalid token")
}
