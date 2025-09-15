package auth

import (
	"gallery/backend/internal/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Refresh(c *gin.Context) {
	rawRefreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error:": "No refresh token in cookies"})
		return
	}

	claims, err := h.authService.ParseJWT(rawRefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	accessToken, err := h.authService.GenerateJWT(claims.Subject) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	refreshToken, err := h.authService.GenerateRefreshJWT(claims.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	setTokenToCookies(c, refreshToken)
	c.JSON(http.StatusOK, types.AuthResponse{
		Token: accessToken,
		Expired: int(time.Minute * 10),
	})
}