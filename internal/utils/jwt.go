package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nunusavi/task-manager/internal/middleware"
)

var jwtSecret = []byte("your_secret_key")

func GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token valid for 72 hours
	}
	_, tokenString, err := middleware.TokenAuth.Encode(claims)
	return tokenString, err
}
