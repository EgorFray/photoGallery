package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	aService "gallery/backend/internal/service/auth"
)


func Authorization(a aService.AuthServiceInterface, c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	claims, err := a.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		c.Abort()
		return
	}

	if claims.Subject == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subject of JWT payload can't be empty"})
		c.Abort()
		return
	}

	c.Set("userID", claims.Subject)

	c.Next()
}