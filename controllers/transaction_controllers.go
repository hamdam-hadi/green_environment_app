package controllers

import (
	"green_environment_app/models"
	"green_environment_app/services"
	"green_environment_app/utils"
	"net/http"
	"strconv"

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
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}
	transaction, err := tc.TransactionService.GetTransactionByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Transaction not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Transaction details fetched successfully",
		"transaction": transaction,
	})
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	if err := tc.TransactionService.CreateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create transaction"})
	}
	return c.JSON(http.StatusCreated, "Transaction created successfully")
}

func (tc *TransactionController) UpdateTransaction(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}

	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	transaction.ID = uint(id)
	if err := tc.TransactionService.UpdateTransaction(&transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update transaction"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction updated successfully"})
}

func (tc *TransactionController) DeleteTransaction(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid transaction ID"})
	}
	if err := tc.TransactionService.DeleteTransaction(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete transaction"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Transaction deleted successfully"})
}
