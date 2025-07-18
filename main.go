package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	ID int `json: "id"`
	Image string `json: "image"`
	Description string `json: "description"`
}


func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	errHandle(err)
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var pst Post
		if err := rows.Scan(&pst.ID, &pst.Image, &pst.Description); err != nil {
			return posts, err
		}
		posts = append(posts, pst)
	}
	if err = rows.Err(); err != nil {
		return posts, err
	}
	return posts, nil
}


func main() {
	connStr := "user=admin dbname=galery sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	errHandle(err)

	err = db.Ping()
	errHandle(err)

	getPosts(db)
}