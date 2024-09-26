package repositories_test

import (
	"green_environment_app/mocks"
	"green_environment_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindChallengeByID(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	challengeID := uint(1)
	expectedChallenge := &models.Challenge{
		ID:    challengeID,
		Title: "Plant 10 Trees",
	}

	mockChallengeRepo.On("FindByID", challengeID).Return(expectedChallenge, nil)

	challenge, err := mockChallengeRepo.FindByID(challengeID)

	assert.NoError(t, err)
	assert.Equal(t, expectedChallenge, challenge)
	mockChallengeRepo.AssertExpectations(t)
}

func TestSaveChallenge(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	newChallenge := &models.Challenge{
		Title: "Clean up the Beach",
	}

	mockChallengeRepo.On("Save", newChallenge).Return(nil)

	err := mockChallengeRepo.Save(newChallenge)

	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}

func TestFindAllChallenges(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)

	expectedChallenges := []models.Challenge{
		{ID: 1, Title: "Plant 10 Trees"},
		{ID: 2, Title: "Clean up the Beach"},
	}

	mockChallengeRepo.On("FindAll").Return(expectedChallenges, nil)

	challenges, err := mockChallengeRepo.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedChallenges, challenges)
	mockChallengeRepo.AssertExpectations(t)
}

func TestUpdateChallenge(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)

	updatedChallenge := &models.Challenge{
		ID:    1,
		Title: "Updated Challenge",
	}

	mockChallengeRepo.On("Update", updatedChallenge).Return(nil)

	err := mockChallengeRepo.Update(updatedChallenge)

	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}

func TestDeleteChallenge(t *testing.T) {
	mockChallengeRepo := new(mocks.ChallengeRepository)
	challengeID := uint(1)

	mockChallengeRepo.On("Delete", challengeID).Return(nil)

	err := mockChallengeRepo.Delete(challengeID)

	assert.NoError(t, err)
	mockChallengeRepo.AssertExpectations(t)
}
