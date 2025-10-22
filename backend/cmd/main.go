package main

import (
	"database/sql"

	cfg "gallery/backend/config"
	"gallery/backend/internal/handlers/middleware"
	PostsHandlers "gallery/backend/internal/handlers/posts"
	postsRepo "gallery/backend/internal/repository/posts"
	postsService "gallery/backend/internal/service/posts"

	userHandlers "gallery/backend/internal/handlers/user"
	userRepo "gallery/backend/internal/repository/user"
	userService "gallery/backend/internal/service/user"

	authHandlers "gallery/backend/internal/handlers/auth"
	authService "gallery/backend/internal/service/auth"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	config := cfg.InitConfig()
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

	// Auth
	authSvc := authService.NewAuthService(config)
	authHandlers := authHandlers.NewAuthHandler(userSvc, authSvc)

	router := gin.Default()

	router.Static("/postsImg", "./images/postsImg")
	router.Static("/avatars", "./images/avatars")
	router.Use(cors.New(cfg.CorsConfig()))
	// post routers
	protected := router.Group("/")
	protected.Use(middleware.Authorization(authSvc)) 
	{
		protected.GET("/posts", postsHandlers.GetPosts)
		protected.GET("/posts/:id", postsHandlers.GetPostById)
		protected.GET("/posts/search", postsHandlers.SearchPosts)
		protected.POST("/posts", postsHandlers.CreatePost)
		protected.DELETE("/posts/:id", postsHandlers.DeletePost)
		protected.PATCH("user/update", userHandlers.UpdateUser)
	}

	// user routers
	router.POST("/user/create", userHandlers.CreateUser)
	// login
	router.POST("/auth/login", authHandlers.Auth)
	router.POST("/auth/refresh", authHandlers.Refresh)
	router.Run(":8081")
}