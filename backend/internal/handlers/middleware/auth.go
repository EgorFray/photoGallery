package middleware

import (
	"net/http"
	"strings"
	"time"

	aService "gallery/backend/internal/service/auth"

	"github.com/gin-gonic/gin"
)


func Authorization(a aService.AuthServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
				return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.ParseJWT(tokenString)
		if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
				return
		}

		if claims.ExpiresAt < time.Now().Unix() {
    	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
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
}