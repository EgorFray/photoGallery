package user

import (
	"fmt"
	service "gallery/backend/internal/service/user"
	"gallery/backend/internal/types"
	"gallery/backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req types.UserRequest

	if err := c.ShouldBind(&req); err != nil {
		fmt.Println("Request:", req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("HERE IS THE PROBLEM", err)
      return
	}

	if req.Avatar == "" {
		req.Avatar = "/avatars/default-icon.png"
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userId, err := u.userService.CreateUser(req, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, userId)
}