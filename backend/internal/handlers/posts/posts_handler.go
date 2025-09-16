package posts

import (
	"database/sql"
	"gallery/backend/internal/service/posts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService posts.PostServiceInterface
}

func NewPostHandler(postService posts.PostServiceInterface) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no userID"})
	}

	posts, err := h.postService.GetPosts(userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no userID"})
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := h.postService.GetPostById(id, userId.(string))

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	} else if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is no userID"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	description := c.PostForm("description")
	if description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	post, err := h.postService.CreatePost(file, description, userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, post)
}

func (h *PostHandler) SearchPosts(c *gin.Context) {
	queryUrl := c.Query("description")
	posts, err := h.postService.SearchPosts(queryUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	err = h.postService.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}