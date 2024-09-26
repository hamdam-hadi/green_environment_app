package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

type UserService interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	return s.repo.Save(user)
}

func (s *userService) Login(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}
