package repository

import "database/sql"

type Post struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
}

var db *sql.DB


func DbCallGetPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, image, description FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var pst Post
		err := rows.Scan(&pst.ID, &pst.Image, &pst.Description)
		if err != nil {
			return nil, err
		}
		posts = append(posts, pst)		
		}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, err
}