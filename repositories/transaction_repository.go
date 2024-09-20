package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

// TransactionRepository defines methods to interact with the database
type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}

// transactionRepository implements TransactionRepository
type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// Save creates a new transaction record
func (r *transactionRepository) Save(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// FindByID fetches a transaction by its ID
func (r *transactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Update modifies an existing transaction record
func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// Delete removes a transaction by its ID
func (r *transactionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}
