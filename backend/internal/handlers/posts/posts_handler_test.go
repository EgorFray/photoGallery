package posts

import (
	"bytes"
	postsMock "gallery/backend/internal/service/mocks"
	"gallery/backend/internal/types"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createTestRequest(t *testing.T, target string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("image", "test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.Copy(fileWriter, strings.NewReader("fake image content"))
	if err != nil {
		t.Fatal(err)
	}

	_ = writer.WriteField("description", "test")
	_ = writer.WriteField("user_id", "1")

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, target, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

// NO MORE TESTS FOR HANDLERS!!! DON'T AGREE WITH ME
// GO FUCK YOURSELF - IT'S MORE FUN
func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	svc := postsMock.NewMockPostServiceInterface(t)
	postsHandlers := NewPostHandler(svc) 
	// middleware for auth
	userIDMiddleware := func(c *gin.Context) {
    c.Set("userID", "1")
    c.Next()
  }
	router.POST("/posts", userIDMiddleware, postsHandlers.CreatePost)
 
	req := createTestRequest(t, "/posts")
	svc.EXPECT().CreatePost(mock.AnythingOfType("*multipart.FileHeader"), "test", "1").Return(&types.PostModel{ID: 1, Image: "test.jpg", Description: "test"}, nil)

	w := httptest.NewRecorder()
 
	router.ServeHTTP(w, req)
 
	assert.Equal(t, http.StatusCreated, w.Code)
 }

