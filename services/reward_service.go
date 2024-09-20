package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

// RewardService defines methods to interact with reward-related operations
type RewardService interface {
	CreateReward(reward *models.Reward) error
	GetAllRewards() ([]models.Reward, error)
	GetRewardByID(id uint) (*models.Reward, error)
	UpdateReward(reward *models.Reward) error
	DeleteReward(id uint) error
}

// rewardService implements RewardService interface
type rewardService struct {
	repo repositories.RewardRepository
}

// NewRewardService creates a new instance of RewardService
func NewRewardService(repo repositories.RewardRepository) RewardService {
	return &rewardService{repo: repo}
}

func (s *rewardService) CreateReward(reward *models.Reward) error {
	return s.repo.Save(reward)
}

func (s *rewardService) GetAllRewards() ([]models.Reward, error) {
	return s.repo.FindAll()
}

func (s *rewardService) GetRewardByID(id uint) (*models.Reward, error) {
	return s.repo.FindByID(id)
}

func (s *rewardService) UpdateReward(reward *models.Reward) error {
	return s.repo.Update(reward)
}

func (s *rewardService) DeleteReward(id uint) error {
	return s.repo.Delete(id)
}
