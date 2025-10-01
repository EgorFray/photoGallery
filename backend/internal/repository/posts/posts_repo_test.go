package repository_test

import (
	"gallery/backend/internal/testutils"
	"gallery/backend/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDbCallCreatePost(t *testing.T) {
	repo := testutils.SetupTestRepo(t)

	insertedID, err := repo.PostRepo.DbCallCreatePost("/images/test-img.jpg", "test", "1")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	row := repo.DB.QueryRow("SELECT description FROM posts WHERE id = $1", insertedID)
	var desc string
	err = row.Scan(&desc)
	assert.NoError(t, err)
	assert.Equal(t, "test", desc)
}

func TestDbCallGetPostById(t *testing.T) {
	repo := testutils.SetupTestRepo(t)

	insertedID, _ := repo.PostRepo.DbCallCreatePost("/images/test-img.jpg", "test", "1")
	post, err := repo.PostRepo.DbCallGetPostById(int(insertedID), "1")
	assert.NoError(t, err)
	// DbCallGetPostById return types.PostDeetailModel, in which I have timestamp.
	// to keep it simple and to avid mocking timestamp, I just check if fileds assert
	assert.Equal(t, "/images/test-img.jpg", post.Image)
	assert.Equal(t, "test", post.Description)
	assert.False(t, post.CreatedAt.IsZero())
}

func TestDbCallGetPosts(t *testing.T) {
	repo := testutils.SetupTestRepo(t)

	insertedID, err := repo.PostRepo.DbCallCreatePost("/images/test-img.jpg", "test", "1")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	expected := []types.PostModel{
		{
			ID: 1,
			Image: "/images/test-img.jpg",
			Description: "test",
		},
	} 

	posts, err := repo.PostRepo.DbCallGetPosts("1")
	assert.NoError(t, err)
	assert.Equal(t, expected, posts)
}

func TestDbCallSearchPosts(t *testing.T) {
	repo := testutils.SetupTestRepo(t)

	insertedID, err := repo.PostRepo.DbCallCreatePost("/images/testing.jpg", "This is test description for the post", "1")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	expected := []types.PostModel{
		{
			ID: 1,
			Image: "/images/testing.jpg",
			Description: "This is test description for the post",
		},
	} 

	posts, err := repo.PostRepo.DbCallSearchPosts("test", "1")
	assert.NoError(t, err)
	assert.Equal(t, expected, posts)
}