package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)


func setupTestDB(t *testing.T) *sql.DB {
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

	return db

}

func TestDbCallCreatePost(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	insertedID, err := repo.DbCallCreatePost("/images/test-img.jpg", "test", "1")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	row := db.QueryRow("SELECT description FROM posts WHERE id = ?", insertedID)
	var desc string
	err = row.Scan(&desc)
	assert.NoError(t, err)
	assert.Equal(t, "test", desc)
}