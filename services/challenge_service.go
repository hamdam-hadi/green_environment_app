package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

type ChallengeService interface {
	CreateChallenge(challenge *models.Challenge) error
	GetAllChallenges() ([]models.Challenge, error)
	GetChallengeByID(id uint) (*models.Challenge, error)
	UpdateChallenge(challenge *models.Challenge) error
	DeleteChallenge(id uint) error
}

type challengeService struct {
	repo repositories.ChallengeRepository
}

func NewChallengeService(repo repositories.ChallengeRepository) ChallengeService {
	return &challengeService{repo: repo}
}

func (s *challengeService) CreateChallenge(challenge *models.Challenge) error {
	return s.repo.Save(challenge)
}

func (s *challengeService) GetAllChallenges() ([]models.Challenge, error) {
	return s.repo.FindAll()
}

func (s *challengeService) GetChallengeByID(id uint) (*models.Challenge, error) {
	return s.repo.FindByID(id)
}

func (s *challengeService) UpdateChallenge(challenge *models.Challenge) error {
	return s.repo.Update(challenge)
}

func (s *challengeService) DeleteChallenge(id uint) error {
	return s.repo.Delete(id)
}
