package auth

import (
	aService "gallery/backend/internal/service/auth"
	uService "gallery/backend/internal/service/user"
	"gallery/backend/internal/types"
	"net/http"
	"time"

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

	refreshToken, err := h.authService.GenerateRefreshJWT(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	setTokenToCookies(c, refreshToken)
	c.JSON(http.StatusOK, types.AuthResponse{
		Token: accessToken,
		Expired: int(time.Minute * 10),
		User: types.UserResponse{
			ID: userData.ID,
			Name: userData.Name,
			Email: userData.Email,
			Avatar: userData.Avatar,
		},
	})
}

func setTokenToCookies(c *gin.Context, refreshToken string) {
	c.SetCookie("refreshToken", refreshToken, int(time.Hour * 24 * 30), "/", "localhost", false, true)
}