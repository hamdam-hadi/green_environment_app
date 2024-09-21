// user_controller.go
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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "couldn't hash password"})
	}
	user.Password = hashedPassword

	err = uc.UserService.Register(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to register a user"})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registered successfully", "user": user,
	})
}

func (uc *UserController) LoginUser(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	user, err := uc.UserService.Login(loginRequest.Email)
	if err != nil || !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid email or password"})
	}

	// Use utils.JWTOptions directly instead of models.JWTOptions
	jwtOptions := utils.JWTOptions{
		SecretKey:       utils.GetConfig("SECRET_KEY"),
		ExpiresDuration: 72, // Set token expiration duration in hours
	}

	// Generate the JWT using the utils package
	token, err := utils.GenerateJWT(int(user.ID), jwtOptions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "couldn't generate token"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"token":   token,
		"user":    user,
	})
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

func (uc *UserController) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := uc.UserService.Login(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "invalid email address"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User details found successfully", "user": user,
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
