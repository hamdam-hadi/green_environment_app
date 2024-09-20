package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

type InformationRepository interface {
	Create(info *models.Information) error
	GetAll() ([]models.Information, error)
}

type informationRepository struct {
	db *gorm.DB
}

func NewInformationRepository(db *gorm.DB) InformationRepository {
	return &informationRepository{db}
}

func (r *informationRepository) Create(info *models.Information) error {
	return r.db.Create(info).Error
}

func (r *informationRepository) GetAll() ([]models.Information, error) {
	var information []models.Information
	err := r.db.Find(&information).Error
	return information, err
}
