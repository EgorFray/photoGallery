package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type AuthService struct {}

var jwtSecret = []byte("super-secret-key")

func GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}


}