package utils

import (
	"errors"
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

func VerifyToken(inComingToken, tokenType string) error {
	secret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")

	if tokenType == "refresh" {
		secret = os.Getenv("JWT_REFRESH_TOKEN_SECRET")

	}
	token, err := jwt.Parse(inComingToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Invalid Token!!!")
	}
	return nil
}
