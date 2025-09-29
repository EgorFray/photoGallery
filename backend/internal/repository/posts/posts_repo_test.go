package repository_test

import (
	"gallery/backend/internal/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)


func TestDbCallCreatePost(t *testing.T) {
	repo := testutils.SetupTestRepo()

	insertedID, err := repo.PostRepo.DbCallCreatePost("/images/test-img.jpg", "test", "1")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	row := repo.DB.QueryRow("SELECT description FROM posts WHERE id = ?", insertedID)
	var desc string
	err = row.Scan(&desc)
	assert.NoError(t, err)
	assert.Equal(t, "test", desc)
}

func TestDbCallGetPostById(t *testing.T) {

}