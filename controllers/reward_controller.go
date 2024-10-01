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

type RewardController struct {
	RewardService services.RewardService
	JWTConfig     middleware.JWTConfig
}

func NewRewardController(rs services.RewardService) *RewardController {
	return &RewardController{
		RewardService: rs,
		JWTConfig: middleware.JWTConfig{
			SigningKey: []byte(os.Getenv("SECRET_KEY")),
		},
	}
}

func (rc *RewardController) GetRewardByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Reward ID is invalid"})

	}
	reward, err := rc.RewardService.GetRewardByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "reward not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Reward details found seccessfully", "reward": reward,
	})

}

func (rc *RewardController) CreateReward(c echo.Context) error {
	var reward models.Reward
	if err := c.Bind(&reward); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}
	err := rc.RewardService.CreateReward(&reward)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create reward"})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Reward successfully created", "reward": reward,
	})
}
func (rc *RewardController) UpdateReward(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid reward ID"})
	}
	var reward models.Reward
	if err := c.Bind(&reward); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	reward.ID = uint(id) // Now this works because Reward has an ID field
	err = rc.RewardService.UpdateReward(&reward)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update reward"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Reward updated successfully"})
}

func (rc *RewardController) DeleteReward(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid reward ID"})
	}
	err = rc.RewardService.DeleteReward(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete reward"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Reward deleted successfully"})
}
