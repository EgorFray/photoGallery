package config

import "github.com/gin-contrib/cors"

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3001", "https://memorize.marian.fyi"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
	}
}