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

	posts, err := h.postService.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := h.postService.GetPostById(id)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	} else if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}