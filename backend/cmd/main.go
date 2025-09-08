package main

import (
	"database/sql"

	"gallery/backend/config"
	PostsHandlers "gallery/backend/internal/handlers/posts"
	postsRepo "gallery/backend/internal/repository/posts"
	postsService "gallery/backend/internal/service/posts"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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


// // User endpoints
// // Later transfer it to utils or something
// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// func (u *UserHandler)createUser(c *gin.Context) {
// 	var req types.UserRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//       return
// 	}

// 	if req.Avatar == "" {
// 		req.Avatar = "/avatars/default-icon.png"
// 	}

// 	hashedPassword, err := hashPassword(req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	userId, err := u.uRepo.DbCallCreateUser(req.Name, req.Email, hashedPassword, req.Avatar)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, userId)
// }

// func login()

func main() {
	config := config.InitConfig()
	var err error
	db, err := sql.Open("postgres", config.PsqlConnUri)
	if err != nil {
		log.Fatal(err)
	}	
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}	

	// Posts
	postRepo := postsRepo.NewPostRepository(db)
	postsSvc := postsService.NewPostService(postRepo)
	postsHandlers := PostsHandlers.NewPostHandler(postsSvc) 

	// uRepo := user.NewUserRepository(db)
	// userHandler := NewUserHandler(uRepo)

	// authSvc := service.NewAuthService(config)

	router := gin.Default()

	router.Static("/postsImg", "./images/postsImg")
	router.Static("/avatars", "./images/avatars")
	router.Use(cors.Default())
	// post routers
	router.GET("/posts", postsHandlers.GetPosts)
	router.GET("/posts/:id", postsHandlers.GetPostById)
	router.GET("/posts/search", postsHandlers.SearchPosts)
	// router.POST("/posts", handler.createPost)
	router.DELETE("/posts/:id", postsHandlers.DeletePost)
	// user routers
	// router.POST("/user/create", userHandler.createUser)
	// login
	// router.POST("/auth/login", authSvc.GenerateJWT)
	router.Run("localhost:8080")
}