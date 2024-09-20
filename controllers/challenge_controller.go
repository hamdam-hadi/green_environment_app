package controllers

import (
	"green_environment_app/services"
	"green_environment_app/utils"
	"net/http"

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
			SigningKey: []byte(utils.GetConfig("SECRET_KEY")),
		},
	}
}

func (cc *ChallengeController) GetChallengeByID(c echo.Context) error {

	return c.JSON(http.StatusOK, "Challenge details")
}

func (cc *ChallengeController) CreateChallenge(c echo.Context) error {

	return c.JSON(http.StatusCreated, "Challenge created")
}

func (cc *ChallengeController) UpdateChallenge(c echo.Context) error {

	return c.JSON(http.StatusOK, "Challenge updated")
}

func (cc *ChallengeController) DeleteChallenge(c echo.Context) error {

	return c.JSON(http.StatusOK, "Challenge deleted")
}
