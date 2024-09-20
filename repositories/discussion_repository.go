package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

// discussion methods
type DiscussionRepository interface {
	Save(discussion *models.Discussion) error
	FindAll() ([]models.Discussion, error)
	FindByID(id uint) (*models.Discussion, error)
	Update(discussion *models.Discussion) error
	Delete(id uint) error
}

type discussionRepository struct {
	db *gorm.DB
}

// creating a new instance discussion repository
func NewDiscussionRepository(db *gorm.DB) DiscussionRepository {
	return &discussionRepository{db}
}

func (r *discussionRepository) Save(discussion *models.Discussion) error {
	return r.db.Create(discussion).Error
}

func (r *discussionRepository) FindAll() ([]models.Discussion, error) {
	var discussions []models.Discussion
	err := r.db.Find(&discussions).Error
	return discussions, err
}

func (r *discussionRepository) FindByID(id uint) (*models.Discussion, error) {
	var discussion models.Discussion
	err := r.db.First(&discussion, id).Error
	return &discussion, err
}

func (r *discussionRepository) Update(discussion *models.Discussion) error {
	return r.db.Save(discussion).Error
}

func (r *discussionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Discussion{}, id).Error
}
