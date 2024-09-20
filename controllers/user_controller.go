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

type UserController struct {
	UserService services.UserService
	JWTConfig   middleware.JWTConfig
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		UserService: us,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(utils.GetConfig("SECRET_KEY")),
		},
	}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	return c.JSON(http.StatusCreated, "User registered")
}

func (uc *UserController) LoginUser(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"emial"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}
	return c.JSON(http.StatusOK, "User logged in")

}

func (uc *UserController) GetUserByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User ID is not valid"})
	}
	user, err := uc.UserService.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User details fetched successfully", "user": user,
	})
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	user.ID = uint(id)
	err = uc.UserService.UpdateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}
	err = uc.UserService.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
