package repositories_test

import (
	"green_environment_app/mocks"
	"green_environment_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindProductByID(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uint(1)
	expectedProduct := &models.Product{
		ID:    productID,
		Name:  "Eco-friendly Bottle",
		Price: 15.99,
	}

	mockProductRepo.On("FindByID", productID).Return(expectedProduct, nil)

	product, err := mockProductRepo.FindByID(productID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockProductRepo.AssertExpectations(t)
}

func TestSaveProduct(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	newProduct := &models.Product{
		Name:  "Reusable Bag",
		Price: 8.99,
	}

	mockProductRepo.On("Save", newProduct).Return(nil)

	err := mockProductRepo.Save(newProduct)

	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}

func TestFindAllProducts(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	expectedProducts := []models.Product{
		{ID: 1, Name: "Eco-friendly Bottle", Price: 15.99},
		{ID: 2, Name: "Reusable Bag", Price: 8.99},
	}

	mockProductRepo.On("FindAll").Return(expectedProducts, nil)

	products, err := mockProductRepo.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	mockProductRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	updatedProduct := &models.Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 12.99,
	}

	mockProductRepo.On("Update", updatedProduct).Return(nil)

	err := mockProductRepo.Update(updatedProduct)

	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uint(1)

	mockProductRepo.On("Delete", productID).Return(nil)

	err := mockProductRepo.Delete(productID)

	assert.NoError(t, err)
	mockProductRepo.AssertExpectations(t)
}
