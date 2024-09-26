package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

type DiscussionService interface {
	CreateDiscussion(discussion *models.Discussion) error
	GetAllDiscussions() ([]models.Discussion, error)
	GetDiscussionByID(id uint) (*models.Discussion, error)
	UpdateDiscussion(discussion *models.Discussion) error
	DeleteDiscussion(id uint) error
}

type discussionService struct {
	repo repositories.DiscussionRepository
}

func NewDiscussionService(repo repositories.DiscussionRepository) DiscussionService {
	return &discussionService{repo: repo}
}

func (s *discussionService) CreateDiscussion(discussion *models.Discussion) error {
	return s.repo.Save(discussion)
}

func (s *discussionService) GetAllDiscussions() ([]models.Discussion, error) {
	return s.repo.FindAll()
}

func (s *discussionService) GetDiscussionByID(id uint) (*models.Discussion, error) {
	return s.repo.FindByID(id)
}

func (s *discussionService) UpdateDiscussion(discussion *models.Discussion) error {
	return s.repo.Update(discussion)
}

func (s *discussionService) DeleteDiscussion(id uint) error {
	return s.repo.Delete(id)
}
