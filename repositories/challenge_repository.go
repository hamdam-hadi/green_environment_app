package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

type ChallengeRepository interface {
	Save(challenge *models.Challenge) error
	FindAll() ([]models.Challenge, error)
	FindByID(id uint) (*models.Challenge, error)
	Update(challenge *models.Challenge) error
	Delete(id uint) error
}

type challengeRepository struct {
	db *gorm.DB
}

func NewChallengeRepository(db *gorm.DB) ChallengeRepository {
	return &challengeRepository{db}
}

func (r *challengeRepository) Save(challenge *models.Challenge) error {
	return r.db.Create(challenge).Error
}

func (r *challengeRepository) FindAll() ([]models.Challenge, error) {
	var challenges []models.Challenge
	err := r.db.Find(&challenges).Error
	return challenges, err
}

func (r *challengeRepository) FindByID(id uint) (*models.Challenge, error) {
	var challenge models.Challenge
	err := r.db.First(&challenge, id).Error
	return &challenge, err
}

func (r *challengeRepository) Update(challenge *models.Challenge) error {
	return r.db.Save(challenge).Error
}

func (r *challengeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Challenge{}, id).Error
}
