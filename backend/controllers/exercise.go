package controllers

import (
	"learnit/config"
	"learnit/models"

	"github.com/gofiber/fiber/v2"
)

// GetExercises récupère tous les exercices disponibles
func GetExercises(c *fiber.Ctx) error {
	var exercises []models.Exercise
	result := config.DB.Preload("InitialFiles").Preload("TestCases").Find(&exercises)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la récupération des exercices",
		})
	}

	return c.JSON(exercises)
}

// GetExercise récupère un exercice spécifique
func GetExercise(c *fiber.Ctx) error {
	id := c.Params("id")
	var exercise models.Exercise
	result := config.DB.Preload("InitialFiles").Preload("TestCases").First(&exercise, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Exercice non trouvé",
		})
	}

	return c.JSON(exercise)
}

// StartExercise démarre un exercice (crée l'environnement Docker)
func StartExercise(c *fiber.Ctx) error {
	id := c.Params("id")
	var exercise models.Exercise
	if err := config.DB.First(&exercise, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Exercice non trouvé",
		})
	}

	// TODO: Implémenter la création de l'environnement Docker

	return c.JSON(fiber.Map{
		"message":  "Environnement créé avec succès",
		"exercise": exercise,
	})
}

// SubmitExercise soumet une solution pour un exercice
func SubmitExercise(c *fiber.Ctx) error {
	type SubmissionData struct {
		Files []struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		} `json:"files"`
	}

	var submission SubmissionData
	if err := c.BodyParser(&submission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	id := c.Params("id")
	var exercise models.Exercise
	if err := config.DB.Preload("TestCases").First(&exercise, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Exercice non trouvé",
		})
	}

	// TODO: Exécuter les tests dans l'environnement Docker
	// Pour l'instant, on simule un succès
	score := 0
	maxScore := 0

	for _, testCase := range exercise.TestCases {
		maxScore += testCase.Points
		// Simulation : on considère que tous les tests passent
		score += testCase.Points
	}

	// Enregistrer le résultat
	userID := c.Locals("user").(models.User).ID
	result := models.UserExerciseResult{
		UserID:     userID,
		ExerciseID: exercise.ID,
		Status:     "completed",
		Score:      score,
		MaxScore:   maxScore,
	}

	config.DB.Create(&result)

	return c.JSON(fiber.Map{
		"score": score,
		"total": maxScore,
	})
}

// CreateExercise crée un nouvel exercice
func CreateExercise(c *fiber.Ctx) error {
	var exercise models.Exercise
	if err := c.BodyParser(&exercise); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	result := config.DB.Create(&exercise)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création de l'exercice",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(exercise)
}

// UpdateExercise met à jour un exercice existant
func UpdateExercise(c *fiber.Ctx) error {
	id := c.Params("id")
	var exercise models.Exercise
	if err := config.DB.First(&exercise, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Exercice non trouvé",
		})
	}

	if err := c.BodyParser(&exercise); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	config.DB.Save(&exercise)
	return c.JSON(exercise)
}

// DeleteExercise supprime un exercice
func DeleteExercise(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Exercise{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la suppression de l'exercice",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
