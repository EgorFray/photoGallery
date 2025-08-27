package main

import (
	"database/sql"
	"gallery/backend/internal/repository"
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

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
} 

var db *sql.DB

// TESTING NEW PROJECT STRUCTURE

// func dbCallGetPosts() ([]Post, error) {
// 	rows, err := db.Query("SELECT id, image, description FROM posts")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var posts []Post

// 	for rows.Next() {
// 		var pst Post
// 		err := rows.Scan(&pst.ID, &pst.Image, &pst.Description)
// 		if err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, pst)		
// 		}
// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return posts, err
// }

// This function is for endpoint GetPostById and open Detail view of post on frontend
func dbCallGetPostById(id int) (PostDetail, error) {
	var post PostDetail
	err := db.QueryRow("SELECT image, description, created_at FROM posts WHERE id = $1", id).Scan(&post.Image, &post.Description, &post.CreatedAt)
	return post, err
}

func dbCallCreatePost(imagePath, description string) (int64, error) {
	var insertedID int64

	err := db.QueryRow("INSERT INTO posts (image, description) VALUES ($1, $2) RETURNING id", imagePath, description).Scan(&insertedID)
	return insertedID, err
}

// This function is for endpoit CreatePost and return created Post which then is added to List of posts on frontend
func dbCallGetCreatedPost(insertedID int64) (Post, error) {
	var post Post
	err := db.QueryRow("SELECT id, image, description FROM posts WHERE id = $1", insertedID).Scan(&post.ID, &post.Image, &post.Description)
	return post, err
}

func (h *Handler) getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
 
	posts, err := h.repo.DbCallGetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.IndentedJSON(http.StatusOK, posts)
	}

func getPostById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := dbCallGetPostById(id)

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

	insertedID, err := dbCallCreatePost(imagePath, description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post, err := dbCallGetCreatedPost(insertedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusCreated, post)
}

func dbCallDeletePost(id int) (error) {
	_, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	return err
}

func deletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	err = dbCallDeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func dbCallSearchPosts(queryUrl string) ([]Post, error) {
		queryDb := "SELECT id, image, description FROM posts WHERE description ILIKE $1"
		rows, err := db.Query(queryDb, "%"+queryUrl+"%")
		if err != nil {
			return nil, err
		}

		defer rows.Close()
	
		var posts []Post
	
		for rows.Next() {
			var pst Post
			err := rows.Scan(&pst.ID, &pst.Image, &pst.Description)
			if err != nil {
				log.Fatal(err)
			}	
			posts = append(posts, pst)		
			}
		err = rows.Err()
		if err != nil {
			return nil, err
		}		
		return posts, err
}
	
func searchPosts(c *gin.Context) {
		queryUrl := c.Query("description")

		posts, err := dbCallSearchPosts(queryUrl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	
		c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	connStr := "user=admin password=admin dbname=galery sslmode=disable host=localhost port=5432"
	var err error
	db, err = sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}	

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}	

	repo := repository.New(db)
	handler := NewHandler(repo)

	router := gin.Default()

	router.Static("/images", "./images")
	router.Use(cors.Default())
	router.GET("/posts", handler.getPosts)
	router.GET("/posts/:id", getPostById)
	router.GET("/posts/search", searchPosts)
	router.POST("/posts", postPosts)
	router.DELETE("/posts/:id", deletePost)
	router.Run("localhost:8080")
}