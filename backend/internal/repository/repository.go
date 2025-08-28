package repository

import (
	"database/sql"
	"log"
	"time"
)


type PostRepository interface {
	DbCallGetPosts() ([]Post, error)
	DbCallGetPostById(id int) (PostDetail, error)
	DbCallCreatePost(imagePath, description string) (int64, error)
	DbCallGetCreatedPost(insertedID int64) (Post, error)
	DbCallDeletePost(id int) (error)
	DbCallSearchPosts(queryUrl string) ([]Post, error)
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Post struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Description string `json:"description"`
}

type PostDetail struct {
	Image string `json:"image"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}


func (r *Repository) DbCallGetPosts() ([]Post, error) {
	rows, err := r.db.Query("SELECT id, image, description FROM posts")
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

// This function is for endpoint GetPostById and open Detail view of post on frontend
func (r *Repository) DbCallGetPostById(id int) (PostDetail, error) {
	var post PostDetail
	err := r.db.QueryRow("SELECT image, description, created_at FROM posts WHERE id = $1", id).Scan(&post.Image, &post.Description, &post.CreatedAt)
	return post, err
}

func (r *Repository) DbCallCreatePost(imagePath, description string) (int64, error) {
	var insertedID int64

	err := r.db.QueryRow("INSERT INTO posts (image, description) VALUES ($1, $2) RETURNING id", imagePath, description).Scan(&insertedID)
	return insertedID, err
}

// This function is for endpoit CreatePost and return created Post which then is added to List of posts on frontend
func (r *Repository) DbCallGetCreatedPost(insertedID int64) (Post, error) {
	var post Post
	err := r.db.QueryRow("SELECT id, image, description FROM posts WHERE id = $1", insertedID).Scan(&post.ID, &post.Image, &post.Description)
	return post, err
}

func (r *Repository) DbCallDeletePost(id int) (error) {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = $1", id)
	return err
}

func (r *Repository) DbCallSearchPosts(queryUrl string) ([]Post, error) {
	queryDb := "SELECT id, image, description FROM posts WHERE description ILIKE $1"
	rows, err := r.db.Query(queryDb, "%"+queryUrl+"%")
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