package auth

import (
	aService "gallery/backend/internal/service/auth"
	uService "gallery/backend/internal/service/user"
	"gallery/backend/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService uService.UserServiceInterface
	authService aService.AuthServiceInterface
}

func NewAuthHandler(userService uService.UserServiceInterface, authService aService.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{userService: userService, authService: authService}
}

func (h *AuthHandler) Auth(c *gin.Context) {
	var authData types.AuthRequst

	if err := c.ShouldBindJSON(&authData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
	}

	userData, err := h.userService.GetUserByEmail(authData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !aService.CheckPasswordHash(userData.Password, authData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auth failed. Please check your credentials and try again"})
		return
	}

	accessToken, err := h.authService.GenerateJWT(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
}