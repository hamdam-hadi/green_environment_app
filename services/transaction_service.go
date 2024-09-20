package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

// TransactionService defines methods to interact with transaction-related operations
type TransactionService interface {
	CreateTransaction(transaction *models.Transaction) error
	GetAllTransactions() ([]models.Transaction, error)
}

// transactionService implements TransactionService interface
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

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.FindAll()
}
