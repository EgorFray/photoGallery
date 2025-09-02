package main

import (
	"database/sql"
	"gallery/backend/internal/repository"
	"gallery/backend/internal/types"
	"gallery/backend/internal/user"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
} 

type UserHandler struct {
	uRepo *user.UserRepository
}

func NewUserHandler(uRepo *user.UserRepository) *UserHandler {
	return &UserHandler{uRepo: uRepo}
}

var db *sql.DB


func (h *Handler) getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
 
	posts, err := h.repo.DbCallGetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.IndentedJSON(http.StatusOK, posts)
	}

func (h *Handler) getPostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := h.repo.DbCallGetPostById(id)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	} else if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
	}

func (h *Handler) createPost(c *gin.Context) {
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

	filePath := filepath.Join("postsImg", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image file: " + err.Error()})
		return
	}

	imagePath := "/" + filePath

	insertedID, err := h.repo.DbCallCreatePost(imagePath, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post, err := h.repo.DbCallGetCreatedPost(insertedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusCreated, post)
}

func (h *Handler) deletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	err = h.repo.DbCallDeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
	
func (h *Handler) searchPosts(c *gin.Context) {
		queryUrl := c.Query("description")

		posts, err := h.repo.DbCallSearchPosts(queryUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	
		c.IndentedJSON(http.StatusOK, posts)
}

// User endpoints
// Later transfer it to utils or something
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserHandler)createUser(c *gin.Context) {
	var req types.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
	}

	if req.Avatar == "" {
		req.Avatar = "/avatars/default-icon.png"
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userId, err := u.uRepo.DbCallCreateUser(req.Name, req.Email, hashedPassword, req.Avatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, userId)
}

func login()

func main() {
	connStr := "user=admin password=admin dbname=galery sslmode=disable host=localhost port=5432"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}	
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}	

	repo := repository.New(db)
	uRepo := user.NewUserRepository(db)
	handler := NewHandler(repo)
	userHandler := NewUserHandler(uRepo)

	router := gin.Default()

	router.Static("/postsImg", "images/postsImg")
	router.Static("/avatars", "images/avatars")
	router.Use(cors.Default())
	// post routers
	router.GET("/posts", handler.getPosts)
	router.GET("/posts/:id", handler.getPostById)
	router.GET("/posts/search", handler.searchPosts)
	router.POST("/posts", handler.createPost)
	router.DELETE("/posts/:id", handler.deletePost)
	// user routers
	router.POST("/user/create", userHandler.createUser)
	router.Run("localhost:8080")
}