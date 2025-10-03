package posts

import (
	postsSvc "gallery/backend/internal/service"
	"gallery/backend/internal/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	svc := postsSvc.NewMockPostServiceInterface(t)
	postsHandlers := NewPostHandler(svc) 

	svc.EXPECT().CreatePost().Return(&types.PostModel{ID: "1", Image: "test.jpg", Description: "test"}, nil)

	router.POST("/posts", postsHandlers.CreatePost)
 
  
	req, _ := createMultipartRequest(t, "/posts")
	w := httptest.NewRecorder()
 
 
	router.ServeHTTP(w, req)
 
	assert.Equal(t, http.StatusCreated, w.Code)
 }

