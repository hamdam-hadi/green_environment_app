package controllers

import (
	"green_environment_app/services"
	"green_environment_app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TransactionController struct {
	TransactionService services.TransactionService
	JWTConfig          middleware.JWTConfig
}

func NewTransactionController(ts services.TransactionService) *TransactionController {
	return &TransactionController{
		TransactionService: ts,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(utils.GetConfig("SECRET_KEY")),
		},
	}
}

func (tc *TransactionController) GetTransactionByID(c echo.Context) error {

	return c.JSON(http.StatusOK, "Transaction details")
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {

	return c.JSON(http.StatusCreated, "Transaction created")
}
