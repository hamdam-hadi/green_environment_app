package utils

import (
	"green_environment_app/middleware"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTOptions struct {
	ExpiresDuration int
	SecretKey       string
}

func GenerateJWT(userID int, jwtOptions JWTOptions) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtOptions.ExpiresDuration))))

	claims := &middleware.JWTCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(jwtOptions.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
