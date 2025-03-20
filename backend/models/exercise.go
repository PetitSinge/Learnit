package models

import (
	"time"
)

type Exercise struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Title        string     `json:"title" gorm:"not null"`
	Description  string     `json:"description"`
	Category     string     `json:"category"`
	Difficulty   string     `json:"difficulty" gorm:"default:'medium'"` // easy, medium, hard
	DockerImage  string     `json:"docker_image"`                       // image Docker pour l'environnement
	Instructions string     `json:"instructions"`                       // instructions détaillées
	InitialFiles []File     `json:"initial_files" gorm:"foreignKey:ExerciseID"`
	TestCases    []TestCase `json:"test_cases" gorm:"foreignKey:ExerciseID"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type File struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ExerciseID uint      `json:"exercise_id"`
	Path       string    `json:"path"`        // chemin relatif dans l'environnement
	Content    string    `json:"content"`     // contenu initial du fichier
	IsReadOnly bool      `json:"is_readonly"` // si le fichier peut être modifié
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TestCase struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ExerciseID uint      `json:"exercise_id"`
	Name       string    `json:"name"`
	Command    string    `json:"command"`  // commande à exécuter
	Expected   string    `json:"expected"` // sortie attendue
	Points     int       `json:"points" gorm:"default:1"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserExerciseResult struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	ExerciseID uint      `json:"exercise_id"`
	Status     string    `json:"status"` // completed, in_progress, failed
	Score      int       `json:"score"`
	MaxScore   int       `json:"max_score"`
	Duration   int       `json:"duration"` // en secondes
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
