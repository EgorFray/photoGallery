package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)


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

func TestDbCallGetPostById(t *testing.T) {

}