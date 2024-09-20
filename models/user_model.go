package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"` // Could be "admin" or "user"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
