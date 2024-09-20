package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Save(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}
