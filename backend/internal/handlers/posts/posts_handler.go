package posts

import (
	"gallery/backend/internal/service/posts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService posts.PostServiceInterface
}

func NewPostHandler(postService posts.PostServiceInterface) *PostHandler {
	return &PostHandler{postService: postService}
}

func(h *PostHandler) GetPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	posts, err := h.postService.GetPosts()
	// Check later how to handle errors in handler
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, posts)
}