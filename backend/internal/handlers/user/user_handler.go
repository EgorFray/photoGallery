package user

import (
	service "gallery/backend/internal/service/user"
	"gallery/backend/internal/types"
	"gallery/backend/internal/utils"
	"log"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userId, err := u.userService.CreateUser(req, hashedPassword, file)
	if err != nil {
		if err.Error() == "this email has already been used" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("Create user error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, userId)
}

func (u *UserHandler) GetUserById(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no userID"})
	}

	user, err := u.userService.GetUserById(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no userID"})
	}

	name := c.PostForm("name")
	password := c.PostForm("password")

	file, _ := c.FormFile("avatar")

	err := u.userService.UpdateUser(userId.(string), name, password, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}