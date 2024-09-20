package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

type RewardRepository interface {
	Save(reward *models.Reward) error
	FindAll() ([]models.Reward, error)
	FindByID(id uint) (*models.Reward, error)
	Update(reward *models.Reward) error
	Delete(id uint) error
}

type rewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) RewardRepository {
	return &rewardRepository{db}
}

func (r *rewardRepository) Save(reward *models.Reward) error {
	return r.db.Create(reward).Error
}

func (r *rewardRepository) FindAll() ([]models.Reward, error) {
	var rewards []models.Reward
	err := r.db.Find(&rewards).Error
	return rewards, err
}

func (r *rewardRepository) FindByID(id uint) (*models.Reward, error) {
	var reward models.Reward
	err := r.db.First(&reward, id).Error
	return &reward, err
}

func (r *rewardRepository) Update(reward *models.Reward) error {
	return r.db.Save(reward).Error
}

func (r *rewardRepository) Delete(id uint) error {
	return r.db.Delete(&models.Reward{}, id).Error
}
