package repository

import (
	"database/sql"
	"gallery/backend/internal/types"
	"log"
)

type PostRepositoryInterface interface {
	DbCallGetPosts(userId string) ([]types.PostModel, error)
	DbCallGetPostById(id int, userId string) (types.PostDetailModel, error)
	DbCallCreatePost(imagePath, description, userId string) (int64, error)
	DbCallGetCreatedPost(insertedID int64, userId string) (types.PostModel, error)
	DbCallDeletePost(id int, userId string) (error)
	DbCallSearchPosts(queryUrl, userId string) ([]types.PostModel, error)
}

type PostRepo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) DbCallGetPosts(userId string) ([]types.PostModel, error) {
	rows, err := r.db.Query("SELECT id, image, description FROM posts WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []types.PostModel

	for rows.Next() {
		var pst types.PostModel
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
func (r *PostRepo) DbCallGetPostById(id int, userId string) (types.PostDetailModel, error) {
	var post types.PostDetailModel
	err := r.db.QueryRow("SELECT image, description, created_at FROM posts WHERE id = $1 AND user_id = $2", id, userId).Scan(&post.Image, &post.Description, &post.CreatedAt)
	return post, err
}

func (r *PostRepo) DbCallCreatePost(imagePath, description, userId string) (int64, error) {
	var insertedID int64

	err := r.db.QueryRow("INSERT INTO posts (image, description, user_id) VALUES ($1, $2, $3) RETURNING id", imagePath, description, userId).Scan(&insertedID)
	return insertedID, err
}

// This function is for endpoit CreatePost and return created Post which then is added to List of posts on frontend
func (r *PostRepo) DbCallGetCreatedPost(insertedID int64, userId string) (types.PostModel, error) {
	var post types.PostModel
	err := r.db.QueryRow("SELECT id, image, description FROM posts WHERE id = $1 AND user_id = $2", insertedID, userId).Scan(&post.ID, &post.Image, &post.Description)
	return post, err
}

func (r *PostRepo) DbCallDeletePost(id int, userId string) (error) {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = $1 AND user_id = $2", id, userId)
	return err
}

func (r *PostRepo) DbCallSearchPosts(queryUrl, userId string) ([]types.PostModel, error) {
	queryDb := "SELECT id, image, description FROM posts WHERE description ILIKE $1 AND user_id = $2"
	rows, err := r.db.Query(queryDb, "%"+queryUrl+"%", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []types.PostModel

	for rows.Next() {
		var pst types.PostModel
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