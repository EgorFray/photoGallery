package main

import (
	"database/sql"

	"gallery/backend/config"
	PostsHandlers "gallery/backend/internal/handlers/posts"
	postsRepo "gallery/backend/internal/repository/posts"
	postsService "gallery/backend/internal/service/posts"

	userHandlers "gallery/backend/internal/handlers/user"
	userRepo "gallery/backend/internal/repository/user"
	userService "gallery/backend/internal/service/user"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	// User
	userRepo := userRepo.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepo)
	userHandlers := userHandlers.NewUserHandler(userSvc)


	router := gin.Default()

	router.Static("/postsImg", "./images/postsImg")
	router.Static("/avatars", "./images/avatars")
	router.Use(cors.Default())
	// post routers
	router.GET("/posts", postsHandlers.GetPosts)
	router.GET("/posts/:id", postsHandlers.GetPostById)
	router.GET("/posts/search", postsHandlers.SearchPosts)
	router.POST("/posts", postsHandlers.CreatePost)
	router.DELETE("/posts/:id", postsHandlers.DeletePost)
	// user routers
	router.POST("/user/create", userHandlers.CreateUser)
	// login
	// router.POST("/auth/login", authSvc.GenerateJWT)
	router.Run("localhost:8080")
}