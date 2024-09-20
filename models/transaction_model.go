package models

import (
	"time"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
