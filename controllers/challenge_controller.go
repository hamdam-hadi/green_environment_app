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

type ChallengeController struct {
	ChallengeService services.ChallengeService
	JWTConfig        middleware.JWTConfig
}

func NewChallengeController(cs services.ChallengeService) *ChallengeController {
	return &ChallengeController{
		ChallengeService: cs,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(os.Getenv("SECRET_KEY")),
		},
	}
}

func (cc *ChallengeController) GetChallengeByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	challenge, err := cc.ChallengeService.GetChallengeByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Challenge not found"})
	}

	return c.JSON(http.StatusOK, challenge)
}

func (cc *ChallengeController) CreateChallenge(c echo.Context) error {
	var challenge models.Challenge
	if err := c.Bind(&challenge); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	if err := cc.ChallengeService.CreateChallenge(&challenge); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create challenge"})
	}

	return c.JSON(http.StatusCreated, challenge)
}

func (cc *ChallengeController) UpdateChallenge(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	var challenge models.Challenge
	if err := c.Bind(&challenge); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	challenge.ID = uint(id)
	if err := cc.ChallengeService.UpdateChallenge(&challenge); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update challenge"})
	}

	return c.JSON(http.StatusOK, challenge)
}

func (cc *ChallengeController) DeleteChallenge(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	if err := cc.ChallengeService.DeleteChallenge(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete challenge"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Challenge deleted successfully"})
}
