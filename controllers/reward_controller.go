package controllers

import (
	"green_environment_app/services"
	"green_environment_app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RewardController struct {
	RewardService services.RewardService
	JWTConfig     middleware.JWTConfig
}

func NewRewardController(rs services.RewardService) *RewardController {
	return &RewardController{
		RewardService: rs,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(utils.GetConfig("SECRET_KEY")),
		},
	}
}

func (rc *RewardController) GetRewardByID(c echo.Context) error {

	return c.JSON(http.StatusOK, "Reward details")
}

func (rc *RewardController) CreateReward(c echo.Context) error {

	return c.JSON(http.StatusCreated, "Reward created")
}

func (rc *RewardController) UpdateReward(c echo.Context) error {

	return c.JSON(http.StatusOK, "Reward updated")
}

func (rc *RewardController) DeleteReward(c echo.Context) error {

	return c.JSON(http.StatusOK, "Reward deleted")
}
