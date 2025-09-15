package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gallery/backend/config"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	GenerateJWT(userId string) (string, error)
	GenerateRefreshJWT(userId string) (string, error)
	ParseJWT(rawToken string) (*jwt.StandardClaims, error)
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
		"exp":     time.Now().Add(a.config.AccessTokenLife).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.config.SecretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString, nil
}

// Yes, I know it's almost the same as GenerateJwt. I will fix it in the future. Maybe :)
func (a *AuthService)GenerateRefreshJWT(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(a.config.RefreshTokenLife).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.config.SecretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString, nil
}

func (a *AuthService) ParseJWT(rawToken string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		rawToken, 
		&jwt.StandardClaims{}, 
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(a.config.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("could not parse this token")
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
