package models

import (
	"time"
)

type Ranking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Score     int       `json:"score"`    // Score total
	Level     int       `json:"level"`    // Niveau actuel
	Position  int       `json:"position"` // Position dans le classement
	Category  string    `json:"category"` // Catégorie spécifique (optionnel)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Badge struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	MinLevel    int       `json:"min_level"` // Niveau minimum requis
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserBadge struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	BadgeID   uint      `json:"badge_id"`
	EarnedAt  time.Time `json:"earned_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
