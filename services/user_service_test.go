package services_test

import (
	"errors"
	"green_environment_app/mocks"
	"green_environment_app/models"
	"green_environment_app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Register_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	user := &models.User{ID: 1, Email: "user@example.com", Username: "testuser"}

	// Define Expectations
	mockUserRepo.On("Save", user).Return(nil)

	// Call the service method
	err := service.Register(user)

	// Assertions
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Register_Failure(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	user := &models.User{ID: 1, Email: "khan@example.com", Username: "testuser"}

	// Define Expectations for Failure
	mockUserRepo.On("Save", user).Return(errors.New("failed to register user"))

	// Call the service method
	err := service.Register(user)

	// Assertions
	assert.Error(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Login_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	mockUser := &models.User{ID: 1, Email: "user@example.com", Username: "testuser"}

	// Define Expectations
	mockUserRepo.On("FindByEmail", "user@example.com").Return(mockUser, nil)

	// Call the service method
	user, err := service.Login("user@example.com")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Login_Failure(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Define Expectations for Failure
	mockUserRepo.On("FindByEmail", "user@example.com").Return(nil, errors.New("user not found"))

	// Call the service method
	user, err := service.Login("user@example.com")

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	mockUser := &models.User{ID: 1, Email: "user@example.com", Username: "testuser"}

	// Define Expectations
	mockUserRepo.On("FindByID", uint(1)).Return(mockUser, nil)

	// Call the service method
	user, err := service.GetUserByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_Failure(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Define Expectations for Failure
	mockUserRepo.On("FindByID", uint(1)).Return(nil, errors.New("user not found"))

	// Call the service method
	user, err := service.GetUserByID(1)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	user := &models.User{ID: 1, Email: "user@example.com", Username: "testuser"}

	// Define Expectations
	mockUserRepo.On("Update", user).Return(nil)

	// Call the service method
	err := service.UpdateUser(user)

	// Assertions
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_Failure(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	service := services.NewUserService(mockUserRepo)

	// Mock Data
	user := &models.User{ID: 1, Email: "user@example.com", Username: "testuser"}

	// Define Expectations for Failure
	mockUserRepo.On("Update", user).Return(errors.New("failed to update user"))

	// Call the service method
	err := service.UpdateUser(user)

	// Assertions
	assert.Error(t, err)
	mockUserRepo.AssertExpectations(t)
}
