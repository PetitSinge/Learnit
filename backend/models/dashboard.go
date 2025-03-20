package models

import (
	"time"
)

type UserProgress struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	CoursesStarted  int       `json:"courses_started"`
	CoursesComplete int       `json:"courses_completed"`
	QuizzesTaken    int       `json:"quizzes_taken"`
	QuizAvgScore    float64   `json:"quiz_avg_score"`
	ExercisesDone   int       `json:"exercises_done"`
	TotalPoints     int       `json:"total_points"`
	StudyTime       int       `json:"study_time"` // en minutes
	LastActive      time.Time `json:"last_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Achievement struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Condition   string    `json:"condition"` // critère pour débloquer
	Points      int       `json:"points"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAchievement struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id"`
	AchievementID uint      `json:"achievement_id"`
	UnlockedAt    time.Time `json:"unlocked_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
