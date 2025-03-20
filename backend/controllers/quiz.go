package controllers

import (
	"learnit/config"
	"learnit/models"

	"github.com/gofiber/fiber/v2"
)

// GetQuizzes récupère tous les quiz disponibles
func GetQuizzes(c *fiber.Ctx) error {
	var quizzes []models.Quiz
	result := config.DB.Preload("Questions.Options").Find(&quizzes)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la récupération des quiz",
		})
	}

	return c.JSON(quizzes)
}

// GetQuiz récupère un quiz spécifique
func GetQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	var quiz models.Quiz
	result := config.DB.Preload("Questions.Options").First(&quiz, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Quiz non trouvé",
		})
	}

	return c.JSON(quiz)
}

// SubmitQuiz soumet les réponses d'un quiz
func SubmitQuiz(c *fiber.Ctx) error {
	type Answer struct {
		QuestionID uint `json:"question_id"`
		OptionID   uint `json:"option_id"`
	}

	type SubmissionData struct {
		Answers []Answer `json:"answers"`
	}

	var submission SubmissionData
	if err := c.BodyParser(&submission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	// Calculer le score
	score := 0
	maxScore := 0

	for _, answer := range submission.Answers {
		var option models.Option
		config.DB.First(&option, answer.OptionID)

		var question models.Question
		config.DB.First(&question, answer.QuestionID)

		maxScore += question.Points
		if option.IsCorrect {
			score += question.Points
		}
	}

	// Enregistrer le résultat
	userID := c.Locals("user").(models.User).ID
	var quiz models.Quiz
	if err := config.DB.First(&quiz, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Quiz non trouvé",
		})
	}

	result := models.UserQuizResult{
		UserID:   userID,
		QuizID:   quiz.ID,
		Score:    score,
		MaxScore: maxScore,
	}

	config.DB.Create(&result)

	return c.JSON(fiber.Map{
		"score": score,
		"total": maxScore,
	})
}

// CreateQuiz crée un nouveau quiz
func CreateQuiz(c *fiber.Ctx) error {
	var quiz models.Quiz
	if err := c.BodyParser(&quiz); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	result := config.DB.Create(&quiz)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création du quiz",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(quiz)
}

// UpdateQuiz met à jour un quiz existant
func UpdateQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	var quiz models.Quiz
	if err := config.DB.First(&quiz, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Quiz non trouvé",
		})
	}

	if err := c.BodyParser(&quiz); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	config.DB.Save(&quiz)
	return c.JSON(quiz)
}

// DeleteQuiz supprime un quiz
func DeleteQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Quiz{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la suppression du quiz",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
