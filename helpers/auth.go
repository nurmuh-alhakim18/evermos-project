package helpers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) (string, error) {
	secretKey := GetEnv("APP_SECRET", "")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    GetEnv("APP_NAME", ""),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   strconv.Itoa(userID),
	})

	key := []byte(secretKey)
	return token.SignedString(key)
}

func ValidateJWT(tokenString string) (interface{}, error) {
	secretKey := GetEnv("APP_SECRET", "")
	key := []byte(secretKey)
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Issuer != GetEnv("APP_NAME", "") {
		return nil, errors.New("invalid issuer")
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	return id, nil
}

func GetBearerToken(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
		return "", errors.New("invalid authorization header")
	}

	return parts[1], nil
}
