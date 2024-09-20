package services

import (
	"green_environment_app/models"
	"green_environment_app/repositories"
)

// ProductService defines methods to interact with product-related operations
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

// productService implements ProductService interface
type productService struct {
	repo repositories.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Save(product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
