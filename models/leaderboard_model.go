package models

import (
	"time"
)

type Leaderboard struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	ChallengeID uint      `json:"challenge_id"`
	Rank        int       `json:"rank"`
	Points      int       `json:"points"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
