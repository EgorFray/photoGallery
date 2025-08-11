package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Post struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
}

type PostRequest struct {
	Image string `json:"image"`
	Description string `json:"description"`
}

type PostDetail struct {
	Image string `json:"image"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB


func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, image, description FROM posts")
	errHandle(err)
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var pst Post
		err := rows.Scan(&pst.ID, &pst.Image, &pst.Description)
		errHandle(err)	
		posts = append(posts, pst)		
		}
	err = rows.Err()
	errHandle(err)	

	c.IndentedJSON(http.StatusOK, posts)
	}

func getPostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	var post PostDetail
	err = db.QueryRow("SELECT image, description, created_at FROM posts WHERE id = $1", id).Scan(&post.Image, &post.Description, &post.CreatedAt)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	} else if err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
	}


func postPosts(c *gin.Context) {
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

	filePath := filepath.Join("images", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image file: " + err.Error()})
		return
	}

	imagePath := "/" + filePath
	var insertedID int64

	query := "INSERT INTO posts (image, description) VALUES ($1, $2) RETURNING id"
	err = db.QueryRow(query, imagePath, description).Scan(&insertedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var post Post
	query = "SELECT id, image, description FROM posts WHERE id = $1"
	err = db.QueryRow(query, insertedID).Scan(&post.ID, &post.Image, &post.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusCreated, post)
}

func searchPosts(c *gin.Context) {
	queryDb := "SELECT id, image, description FROM posts WHERE description ILIKE $1"
	queryUrl := c.Query("description")
	rows, err := db.Query(queryDb, "%"+queryUrl+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var pst Post
		err := rows.Scan(&pst.ID, &pst.Image, &pst.Description)
		errHandle(err)	
		posts = append(posts, pst)		
		}
	err = rows.Err()
	errHandle(err)	

	c.IndentedJSON(http.StatusOK, posts)
}
	
func main() {
	connStr := "user=admin password=admin dbname=galery sslmode=disable host=localhost port=5432"
	var err error
	db, err = sql.Open("postgres", connStr)
	defer db.Close()
	errHandle(err)

	err = db.Ping()
	errHandle(err)

	router := gin.Default()

	router.Static("/images", "./images")
	router.Use(cors.Default())
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostById)
	router.GET("/posts/search", searchPosts)
	router.POST("/posts", postPosts)
	router.Run("localhost:8080")
}