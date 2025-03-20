package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"learnit/models"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto-migration des modèles
	err = DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Chapter{},
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.UserQuizResult{},
		&models.Exercise{},
		&models.File{},
		&models.TestCase{},
		&models.UserExerciseResult{},
		&models.UserProgress{},
		&models.Achievement{},
		&models.UserAchievement{},
		&models.Ranking{},
		&models.Badge{},
		&models.UserBadge{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	log.Println("Database connected successfully")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
