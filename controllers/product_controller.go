package controllers

import (
	"green_environment_app/models"
	"green_environment_app/services"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ProductController struct {
	ProductService services.ProductService
	JWTConfig      middleware.JWTConfig
}

func NewProductController(ps services.ProductService) *ProductController {
	return &ProductController{
		ProductService: ps,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(os.Getenv("SECRET_KEY")),
		},
	}
}

func (pc *ProductController) GetProductByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid ID"})
	}
	product, err := pc.ProductService.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "details are invalid"})
	}
	if err := pc.ProductService.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create product"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "product created successfully"})
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"messsage": "invalid product id"})
	}
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}
	product.ID = uint(id)
	if err := pc.ProductService.UpdateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to update product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "product updated successfully"})
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid product ID"})
	}
	if err := pc.ProductService.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to delete product"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "product deleted seccessfully"})
}
