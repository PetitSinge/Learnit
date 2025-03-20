package models

import (
	"time"
)

type Quiz struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	IsTest      bool       `json:"is_test" gorm:"default:false"` // true pour les tests réels
	TimeLimit   int        `json:"time_limit"`                   // en minutes, 0 pour pas de limite
	Questions   []Question `json:"questions" gorm:"foreignKey:QuizID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Question struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	QuizID      uint      `json:"quiz_id"`
	Text        string    `json:"text" gorm:"not null"`
	Options     []Option  `json:"options" gorm:"foreignKey:QuestionID"`
	Explanation string    `json:"explanation"`
	Points      int       `json:"points" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Option struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	QuestionID uint      `json:"question_id"`
	Text       string    `json:"text" gorm:"not null"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserQuizResult struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	QuizID    uint      `json:"quiz_id"`
	Score     int       `json:"score"`
	MaxScore  int       `json:"max_score"`
	Duration  int       `json:"duration"` // en secondes
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
