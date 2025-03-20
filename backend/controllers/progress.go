package controllers

import (
	"learnit/config"
	"learnit/models"

	"github.com/gofiber/fiber/v2"
)

// GetProgress récupère les progrès de l'utilisateur
func GetProgress(c *fiber.Ctx) error {
	userID := c.Locals("user").(models.User).ID

	var progress models.UserProgress
	result := config.DB.Where("user_id = ?", userID).First(&progress)
	if result.Error != nil {
		// Si aucun progrès n'existe, en créer un nouveau
		progress = models.UserProgress{
			UserID: userID,
		}
		config.DB.Create(&progress)
	}

	return c.JSON(progress)
}

// GetAchievements récupère les achievements de l'utilisateur
func GetAchievements(c *fiber.Ctx) error {
	userID := c.Locals("user").(models.User).ID

	var achievements []models.Achievement
	var userAchievements []models.UserAchievement

	config.DB.Find(&achievements)
	config.DB.Where("user_id = ?", userID).Find(&userAchievements)

	type AchievementStatus struct {
		models.Achievement
		Unlocked bool `json:"unlocked"`
	}

	response := make([]AchievementStatus, 0)
	for _, achievement := range achievements {
		unlocked := false
		for _, ua := range userAchievements {
			if ua.AchievementID == achievement.ID {
				unlocked = true
				break
			}
		}
		response = append(response, AchievementStatus{
			Achievement: achievement,
			Unlocked:    unlocked,
		})
	}

	return c.JSON(response)
}

// GetRanking récupère le classement global
func GetRanking(c *fiber.Ctx) error {
	var rankings []models.Ranking
	result := config.DB.Order("score desc").Limit(10).Find(&rankings)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la récupération du classement",
		})
	}

	type RankingWithUser struct {
		models.Ranking
		UserName string `json:"user_name"`
	}

	response := make([]RankingWithUser, 0)
	for _, ranking := range rankings {
		var user models.User
		config.DB.Select("name").First(&user, ranking.UserID)
		response = append(response, RankingWithUser{
			Ranking:  ranking,
			UserName: user.Name,
		})
	}

	return c.JSON(response)
}
