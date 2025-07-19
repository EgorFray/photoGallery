package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Post struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
}

var db *sql.DB


func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM posts")
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
	
func main() {
	connStr := "user=admin dbname=galery sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	defer db.Close()
	errHandle(err)

	err = db.Ping()
	errHandle(err)

	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:8080")

}