package services_test

import (
	"errors"
	"green_environment_app/mocks"
	"green_environment_app/models"
	"green_environment_app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRewardService_GetAllRewards_Success(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	mockRewards := []models.Reward{
		{ID: 1, Name: "Reward 1"},
		{ID: 2, Name: "Reward 2"},
	}

	// Define Expectations
	mockRewardRepo.On("FindAll").Return(mockRewards, nil)

	// Call the service method
	rewards, err := service.GetAllRewards()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockRewards, rewards)

	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_GetAllRewards_Failure(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Define Expectations for Failure
	mockRewardRepo.On("FindAll").Return(nil, errors.New("failed to fetch rewards"))

	// Call the service method
	rewards, err := service.GetAllRewards()

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, rewards)

	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_CreateReward_Success(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	reward := &models.Reward{ID: 1, Name: "New Reward"}

	// Define Expectations
	mockRewardRepo.On("Save", reward).Return(nil)

	// Call the service method
	err := service.CreateReward(reward)

	// Assertions
	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_CreateReward_Failure(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	reward := &models.Reward{ID: 1, Name: "New Reward"}

	// Define Expectations for Failure
	mockRewardRepo.On("Save", reward).Return(errors.New("failed to create reward"))

	// Call the service method
	err := service.CreateReward(reward)

	// Assertions
	assert.Error(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_GetRewardByID_Success(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	mockReward := &models.Reward{ID: 1, Name: "Reward 1"}

	// Define Expectations
	mockRewardRepo.On("FindByID", uint(1)).Return(mockReward, nil)

	// Call the service method
	reward, err := service.GetRewardByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockReward, reward)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_GetRewardByID_Failure(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Define Expectations for Failure
	mockRewardRepo.On("FindByID", uint(1)).Return(nil, errors.New("reward not found"))

	// Call the service method
	reward, err := service.GetRewardByID(1)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, reward)

	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_UpdateReward_Success(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	reward := &models.Reward{ID: 1, Name: "Updated Reward"}

	// Define Expectations
	mockRewardRepo.On("Update", reward).Return(nil)

	// Call the service method
	err := service.UpdateReward(reward)

	// Assertions
	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_UpdateReward_Failure(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Mock Data
	reward := &models.Reward{ID: 1, Name: "Updated Reward"}

	// Define Expectations for Failure
	mockRewardRepo.On("Update", reward).Return(errors.New("failed to update reward"))

	// Call the service method
	err := service.UpdateReward(reward)

	// Assertions
	assert.Error(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_DeleteReward_Success(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Define Expectations
	mockRewardRepo.On("Delete", uint(1)).Return(nil)

	// Call the service method
	err := service.DeleteReward(1)

	// Assertions
	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestRewardService_DeleteReward_Failure(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	service := services.NewRewardService(mockRewardRepo)

	// Define Expectations for Failure
	mockRewardRepo.On("Delete", uint(1)).Return(errors.New("failed to delete reward"))

	// Call the service method
	err := service.DeleteReward(1)

	// Assertions
	assert.Error(t, err)
	mockRewardRepo.AssertExpectations(t)
}
