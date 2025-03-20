package controllers

import (
	"learnit/config"
	"learnit/models"

	"github.com/gofiber/fiber/v2"
)

// GetCourses récupère tous les cours
func GetCourses(c *fiber.Ctx) error {
	var courses []models.Course
	if err := config.DB.Preload("Chapters").Find(&courses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la récupération des cours",
		})
	}
	return c.JSON(courses)
}

// GetCourse récupère un cours par son ID
func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course
	if err := config.DB.Preload("Chapters").First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cours non trouvé",
		})
	}
	return c.JSON(course)
}

// CreateCourse crée un nouveau cours
func CreateCourse(c *fiber.Ctx) error {
	course := new(models.Course)
	if err := c.BodyParser(course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	if err := config.DB.Create(&course).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création du cours",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(course)
}

// UpdateCourse met à jour un cours
func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	course := new(models.Course)

	if err := config.DB.First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cours non trouvé",
		})
	}

	if err := c.BodyParser(course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	if err := config.DB.Save(&course).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la mise à jour du cours",
		})
	}

	return c.JSON(course)
}

// DeleteCourse supprime un cours
func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course

	if err := config.DB.First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cours non trouvé",
		})
	}

	if err := config.DB.Delete(&course).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la suppression du cours",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// AddChapter ajoute un chapitre à un cours
func AddChapter(c *fiber.Ctx) error {
	courseID := c.Params("courseId")
	chapter := new(models.Chapter)

	if err := c.BodyParser(chapter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	// Vérifier si le cours existe
	var course models.Course
	if err := config.DB.First(&course, courseID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cours non trouvé",
		})
	}

	chapter.CourseID = course.ID
	if err := config.DB.Create(&chapter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création du chapitre",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(chapter)
}

// UpdateChapter met à jour un chapitre
func UpdateChapter(c *fiber.Ctx) error {
	id := c.Params("id")
	chapter := new(models.Chapter)

	if err := config.DB.First(&chapter, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Chapitre non trouvé",
		})
	}

	if err := c.BodyParser(chapter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	if err := config.DB.Save(&chapter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la mise à jour du chapitre",
		})
	}

	return c.JSON(chapter)
}

// DeleteChapter supprime un chapitre
func DeleteChapter(c *fiber.Ctx) error {
	id := c.Params("id")
	var chapter models.Chapter

	if err := config.DB.First(&chapter, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Chapitre non trouvé",
		})
	}

	if err := config.DB.Delete(&chapter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la suppression du chapitre",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
