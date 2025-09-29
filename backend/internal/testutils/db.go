package testutils

import (
	"database/sql"
	postRepo "gallery/backend/internal/repository/posts"
	userRepo "gallery/backend/internal/repository/user"
	"log"
)

type TestRepo struct {
	DB *sql.DB
	PostRepo postRepo.PostRepo
	UserRepo userRepo.UserRepository
}

func SetupTestRepo() *TestRepo {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("failed to open sqlite: %v", err)
	} 

	_, err = db.Exec(`
	CREATE TABLE posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		image TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER
	)
	`)

	if err != nil {
		log.Fatalf("failed to create table posts: %v", err)
	}

	return &TestRepo {
		DB: db,
		PostRepo: *postRepo.NewPostRepository(db),
		UserRepo: *userRepo.NewUserRepository(db),
	}
}