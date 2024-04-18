package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(tokenType string, jti string, id string, expire int64) (string, error) {

	accessSecret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	var secret string = accessSecret

	if tokenType == "refresh" {
		secret = refreshSecret
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"jti": jti,
		"exp": expire,
		"iat": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}
