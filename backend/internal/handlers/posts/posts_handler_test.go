package posts

import (
	postsSvc "gallery/backend/internal/service"
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

	svc.EXPECT().CreatePost().Return()

	router.POST("/posts", postsHandlers.CreatePost)
 
	// body - multipart  
	req, _ := http.NewRequest(http.MethodPost, "/posts", {})
 
	w := httptest.NewRecorder()
 
 
	router.ServeHTTP(w, req)
 
	assert.Equal(t, http.StatusCreated, w.Code)
 }

