package service

import (
	"log"
	"time"

	"gallery/backend/config"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
		"sub": userId,
		"exp":     a.config.AccessTokenLife,
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.config.SecretKey)
	if err != nil {
		log.Println(err)
	}
	return tokenString, nil
}

// Yes, I know it's almost the same as GenerateJwt. I will fix it in the future. Maybe :)
func (a *AuthService)GenerateRefreshJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": a.config.RefreshTokenLife,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.config.SecretKey)
	if err != nil {
		log.Println(err)
	}
	return tokenString, nil
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
