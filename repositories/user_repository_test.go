package repositories_test

import (
	"green_environment_app/mocks"
	"green_environment_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	userID := uint(1)
	expectedUser := &models.User{
		ID:       userID,
		Username: "hamdam",
		Email:    "hamdam@example.com",
	}
	mockUserRepo.On("FindByID", userID).Return(expectedUser, nil)
	user, err := mockUserRepo.FindByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestSave(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	newUser := &models.User{
		Username: "newuser",
		Email:    "newuser@example.com",
		Password: "password",
	}

	mockUserRepo.On("Save", newUser).Return(nil)

	err := mockUserRepo.Save(newUser)

	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestFindAll(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	expectedUser := []models.User{
		{ID: 1, Username: "user1", Email: "user1@mail.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}
	mockUserRepo.On("FindAll").Return(expectedUser, nil)
	users, err := mockUserRepo.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, users)
	mockUserRepo.AssertExpectations(t)
}

func TestFindByEmail(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	email := "hamdam@mail.com"
	expectedUser := &models.User{
		ID:       1,
		Username: "hamdam",
		Email:    email,
	}
	mockUserRepo.On("FindByEmail", email).Return(expectedUser, nil)
	user, err := mockUserRepo.FindByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockUserRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	updatedUser := &models.User{
		ID:       1,
		Username: "updateduser",
		Email:    "updateduser@example.com",
	}

	mockUserRepo.On("Update", updatedUser).Return(nil)

	err := mockUserRepo.Update(updatedUser)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userID := uint(1)

	mockUserRepo.On("Delete", userID).Return(nil)

	err := mockUserRepo.Delete(userID)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}
