package repository

import (
	"database/sql"
	"log"
	"testing"
	_ "modernc.org/sqlite"
)


func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("failed to open sqlite", err)
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
		log.Fatalf("failed to creat table posts: %v", err)
	}

	return db

}

func TestDbCallGetPostById(t *testing.T) {
}