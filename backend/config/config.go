package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PsqlConnUri string
	SecretKey string
	AccessTokenLife time.Duration
	RefreshTokenLife time.Duration
}

func InitConfig() *Config {
	LoadEnvVariables()
	return &Config{
		PsqlConnUri: os.Getenv("DATABASE_URL"),
		SecretKey: os.Getenv("JWT_SECRET"),
		AccessTokenLife: time.Minute * 10,
		RefreshTokenLife: time.Hour * 24 * 30,
	}
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}