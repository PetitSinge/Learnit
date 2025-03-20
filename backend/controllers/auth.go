package controllers

import (
	"os"
	"time"

	"learnit/config"
	"learnit/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	input := new(RegisterInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	// Vérifier si l'utilisateur existe déjà
	var existingUser models.User
	result := config.DB.Where("email = ?", input.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Cet email est déjà utilisé",
		})
	}

	// Créer le nouvel utilisateur
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création du compte",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Compte créé avec succès",
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func Login(c *fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Données invalides",
		})
	}

	// Trouver l'utilisateur
	var user models.User
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email ou mot de passe incorrect",
		})
	}

	// Vérifier le mot de passe
	if err := user.ComparePassword(input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Email ou mot de passe incorrect",
		})
	}

	// Créer le token JWT
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expire dans 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erreur lors de la création du token",
		})
	}

	return c.JSON(fiber.Map{
		"token": t,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
