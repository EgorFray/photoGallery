package user_test

import (
	"gallery/backend/internal/testutils"
	"gallery/backend/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetUser(t *testing.T) {
	repo := testutils.SetupTestRepo(t)

	insertedID, err := repo.UserRepo.DbCallCreateUser("TestUser", "test@gmail.com", "qwerty", "test-img.jpg")
	assert.NoError(t, err)
	assert.NotZero(t, insertedID)

	expected := types.UserModel{
		ID:"1", 
		Name:"TestUser", 
		Email:"test@gmail.com", 
		Password:"qwerty", 
		Avatar:"test-img.jpg",
	}

	user, err := repo.UserRepo.DbCallGetUserByEmail("test@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}


