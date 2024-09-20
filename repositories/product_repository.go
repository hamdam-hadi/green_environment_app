package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

// mothods of procudt.
type ProductRepository interface {
	Save(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

// creating instance of new product
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Save(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
