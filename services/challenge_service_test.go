package services_test

import (
	"green_environment_app/mocks"
	"green_environment_app/models"
	"green_environment_app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallengeService_CreateChallenge_Success(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	service := services.NewChallengeService(mockChallengeRepo)

	// Mock Data
	challenge := &models.Challenge{ID: 1, Title: "New Challenge"}

	// Define Expectations
	mockChallengeRepo.On("Save", challenge).Return(nil) // Changed from "Create" to "Save"

	// Call the service method
	err := service.CreateChallenge(challenge)

	// Assertions
	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}

func TestChallengeService_FindAllChallenges(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	service := services.NewChallengeService(mockChallengeRepo)

	// Mock Data
	mockChallenges := []models.Challenge{
		{ID: 1, Title: "Challenge 1"},
		{ID: 2, Title: "Challenge 2"},
	}

	// Define Expectations
	mockChallengeRepo.On("FindAll").Return(mockChallenges, nil)

	// Call the service method
	challenges, err := service.GetAllChallenges()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockChallenges, challenges)

	// Ensure all mock expectations are met
	mockChallengeRepo.AssertExpectations(t)
}

func TestChallengeService_GetChallengeByID_Success(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	service := services.NewChallengeService(mockChallengeRepo)

	// Mock Data
	mockChallenge := &models.Challenge{ID: 1, Title: "Challenge 1"}

	// Define Expectations
	mockChallengeRepo.On("FindByID", uint(1)).Return(mockChallenge, nil)

	// Call the service method
	challenge, err := service.GetChallengeByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockChallenge, challenge)
	mockChallengeRepo.AssertExpectations(t)
}

func TestChallengeService_UpdateChallenge_Success(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	service := services.NewChallengeService(mockChallengeRepo)

	// Mock Data
	challenge := &models.Challenge{ID: 1, Title: "Updated Challenge"}

	// Define Expectations
	mockChallengeRepo.On("Update", challenge).Return(nil)

	// Call the service method
	err := service.UpdateChallenge(challenge)

	// Assertions
	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}

func TestChallengeService_DeleteChallenge_Success(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	service := services.NewChallengeService(mockChallengeRepo)

	// Define Expectations
	mockChallengeRepo.On("Delete", uint(1)).Return(nil)

	// Call the service method
	err := service.DeleteChallenge(1)

	// Assertions
	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}
