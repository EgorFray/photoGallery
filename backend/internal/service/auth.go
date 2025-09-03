package service

import (
	"time"

	"gallery/backend/config"

	"github.com/golang-jwt/jwt"
)

type AuthServiceInterface interface {
	GenerateJWT(userId string) (string, error)
}

type AuthService struct {
	config *config.Config
}

func NewAuthService(config *config.Config) *AuthService {
	return &AuthService{config: config}
}

func (a *AuthService)GenerateJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     a.config.AccessTokenLife,
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.config.SecretKey)
}
