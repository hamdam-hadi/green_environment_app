package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

// TransactionService defines methods to interact with transaction-related operations
type TransactionService interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByID(id uint) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id uint) error
}

// transactionService implements the TransactionService interface
type transactionService struct {
	repo repositories.TransactionRepository
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	return s.repo.Save(transaction)
}

func (s *transactionService) GetTransactionByID(id uint) (*models.Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *transactionService) UpdateTransaction(transaction *models.Transaction) error {
	return s.repo.Update(transaction)
}

func (s *transactionService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}
