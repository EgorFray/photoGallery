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
)

func createMultipartRequest(t *testing.T, target string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", "test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.Copy(fileWriter, strings.NewReader("fake image content"))
	if err != nil {
		t.Fatal(err)
	}

	_ = writer.WriteField("description", "test description")
	_ = writer.WriteField("user_id", "1")

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, target, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	svc := postsMock.NewMockPostServiceInterface(t)
	postsHandlers := NewPostHandler(svc) 

	svc.EXPECT().CreatePost().Return(&types.PostModel{ID: 1, Image: "test.jpg", Description: "test"}, nil)

	router.POST("/posts", postsHandlers.CreatePost)
 
	req := createMultipartRequest(t, "/posts")
	w := httptest.NewRecorder()
 
	router.ServeHTTP(w, req)
 
	assert.Equal(t, http.StatusCreated, w.Code)
 }

