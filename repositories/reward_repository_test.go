package repositories_test

import (
	"green_environment_app/mocks"
	"green_environment_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRewardByID(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	rewardID := uint(1)
	expectedReward := &models.Reward{
		ID:          rewardID,
		Description: "Eco Warrior Badge",
	}

	mockRewardRepo.On("FindByID", rewardID).Return(expectedReward, nil)

	reward, err := mockRewardRepo.FindByID(rewardID)

	assert.NoError(t, err)
	assert.Equal(t, expectedReward, reward)
	mockRewardRepo.AssertExpectations(t)
}

func TestSaveReward(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	newReward := &models.Reward{
		Description: "Tree Planter Badge",
	}

	mockRewardRepo.On("Save", newReward).Return(nil)

	err := mockRewardRepo.Save(newReward)

	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestFindAllRewards(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)

	expectedRewards := []models.Reward{
		{ID: 1, Description: "Eco Warrior Badge"},
		{ID: 2, Description: "Tree Planter Badge"},
	}

	mockRewardRepo.On("FindAll").Return(expectedRewards, nil)

	rewards, err := mockRewardRepo.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedRewards, rewards)
	mockRewardRepo.AssertExpectations(t)
}

func TestUpdateReward(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)

	updatedReward := &models.Reward{
		ID:          1,
		Description: "Updated Badge",
	}

	mockRewardRepo.On("Update", updatedReward).Return(nil)

	err := mockRewardRepo.Update(updatedReward)

	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}

func TestDeleteReward(t *testing.T) {
	mockRewardRepo := new(mocks.RewardRepository)
	rewardID := uint(1)

	mockRewardRepo.On("Delete", rewardID).Return(nil)

	err := mockRewardRepo.Delete(rewardID)

	assert.NoError(t, err)
	mockRewardRepo.AssertExpectations(t)
}
