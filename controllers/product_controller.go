package controllers

import (
	"green_environment_app/services"
	"green_environment_app/utils"
	"net/http"

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
			SigningKey: []byte(utils.GetConfig("SECRET_KEY")),
		},
	}
}

func (pc *ProductController) GetProductByID(c echo.Context) error {

	return c.JSON(http.StatusOK, "Product details")
}

func (pc *ProductController) CreateProduct(c echo.Context) error {

	return c.JSON(http.StatusCreated, "Product created")
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {

	return c.JSON(http.StatusOK, "Product updated")
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {

	return c.JSON(http.StatusOK, "Product deleted")
}
