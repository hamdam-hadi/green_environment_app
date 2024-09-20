package repositories

import (
	"green_environment_app/models"

	"gorm.io/gorm"
)

type LeaderboardRepository interface {
	GetByChallengeID(challengeID uint) ([]models.Leaderboard, error)
}

type leaderboardRepository struct {
	db *gorm.DB
}

func NewLeaderboardRepository(db *gorm.DB) LeaderboardRepository {
	return &leaderboardRepository{db}
}

func (r *leaderboardRepository) GetByChallengeID(challengeID uint) ([]models.Leaderboard, error) {
	var leaderboards []models.Leaderboard
	err := r.db.Where("challenge_id = ?", challengeID).Find(&leaderboards).Error
	return leaderboards, err
}
