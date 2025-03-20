package models

import (
	"time"
)

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	PDFPath     string    `json:"pdf_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Chapters    []Chapter `json:"chapters" gorm:"foreignKey:CourseID"`
}

type Chapter struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CourseID    uint      `json:"course_id"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	PDFPath     string    `json:"pdf_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
