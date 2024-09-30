package services_test

import (
	"errors"
	"green_environment_app/mocks"
	"green_environment_app/models"
	"green_environment_app/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetAllProducts_Success(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Mock Data
	mockProducts := []models.Product{
		{ID: 1, Name: "Product 1"},
		{ID: 2, Name: "Product 2"},
	}

	// Define Expectations
	mockProductRepo.On("FindAll").Return(mockProducts, nil)

	// Call the service method
	products, err := service.GetAllProducts()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockProducts, products)

	mockProductRepo.AssertExpectations(t)
}

func TestProductService_GetAllProducts_Error(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Define Expectations
	mockProductRepo.On("FindAll").Return(nil, errors.New("something went wrong"))

	// Call the service method
	_, err := service.GetAllProducts()

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "something went wrong")

	mockProductRepo.AssertExpectations(t)
}

func TestProductService_CreateProduct_Success(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Mock Data
	product := &models.Product{ID: 1, Name: "New Product"}

	// Define Expectations
	mockProductRepo.On("Save", product).Return(nil)

	// Call the service method
	err := service.CreateProduct(product)

	// Assertions
	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_GetProductByID_Success(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Mock Data
	mockProduct := &models.Product{ID: 1, Name: "Product 1"}

	// Define Expectations
	mockProductRepo.On("FindByID", uint(1)).Return(mockProduct, nil)

	// Call the service method
	product, err := service.GetProductByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockProduct, product)
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_GetProductByID_Error(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Define Expectations
	mockProductRepo.On("FindByID", uint(1)).Return(nil, errors.New("product not found"))

	// Call the service method
	_, err := service.GetProductByID(1)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "product not found")
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_Success(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Mock Data
	product := &models.Product{ID: 1, Name: "Updated Product"}

	// Define Expectations
	mockProductRepo.On("Update", product).Return(nil)

	// Call the service method
	err := service.UpdateProduct(product)

	// Assertions
	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_Error(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Mock Data
	product := &models.Product{ID: 1, Name: "Updated Product"}

	// Define Expectations
	mockProductRepo.On("Update", product).Return(errors.New("failed to update product"))

	// Call the service method
	err := service.UpdateProduct(product)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to update product")
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct_Success(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Define Expectations
	mockProductRepo.On("Delete", uint(1)).Return(nil)

	// Call the service method
	err := service.DeleteProduct(1)

	// Assertions
	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct_Error(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	service := services.NewProductService(mockProductRepo)

	// Define Expectations
	mockProductRepo.On("Delete", uint(1)).Return(errors.New("failed to delete product"))

	// Call the service method
	err := service.DeleteProduct(1)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to delete product")
	mockProductRepo.AssertExpectations(t)
}
