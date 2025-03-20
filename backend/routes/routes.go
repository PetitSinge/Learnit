package routes

import (
	"learnit/controllers"
	"learnit/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API Group
	api := app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	// Courses routes
	courses := api.Group("/courses", middleware.Protected())
	courses.Get("/", controllers.GetCourses)
	courses.Get("/:id", controllers.GetCourse)
	courses.Post("/", middleware.AdminOnly(), controllers.CreateCourse)
	courses.Put("/:id", middleware.AdminOnly(), controllers.UpdateCourse)
	courses.Delete("/:id", middleware.AdminOnly(), controllers.DeleteCourse)

	// Chapters routes
	chapters := api.Group("/chapters", middleware.Protected())
	chapters.Post("/:courseId", middleware.AdminOnly(), controllers.AddChapter)
	chapters.Put("/:id", middleware.AdminOnly(), controllers.UpdateChapter)
	chapters.Delete("/:id", middleware.AdminOnly(), controllers.DeleteChapter)

	// Quiz routes
	quiz := api.Group("/quiz", middleware.Protected())
	quiz.Get("/", controllers.GetQuizzes)
	quiz.Get("/:id", controllers.GetQuiz)
	quiz.Post("/:id/submit", controllers.SubmitQuiz)
	quiz.Post("/", middleware.AdminOnly(), controllers.CreateQuiz)
	quiz.Put("/:id", middleware.AdminOnly(), controllers.UpdateQuiz)
	quiz.Delete("/:id", middleware.AdminOnly(), controllers.DeleteQuiz)

	// Exercise routes
	exercises := api.Group("/exercises", middleware.Protected())
	exercises.Get("/", controllers.GetExercises)
	exercises.Get("/:id", controllers.GetExercise)
	exercises.Post("/:id/start", controllers.StartExercise)
	exercises.Post("/:id/submit", controllers.SubmitExercise)
	exercises.Post("/", middleware.AdminOnly(), controllers.CreateExercise)
	exercises.Put("/:id", middleware.AdminOnly(), controllers.UpdateExercise)
	exercises.Delete("/:id", middleware.AdminOnly(), controllers.DeleteExercise)

	// Progress routes
	progress := api.Group("/progress", middleware.Protected())
	progress.Get("/", controllers.GetProgress)
	progress.Get("/achievements", controllers.GetAchievements)
	progress.Get("/ranking", controllers.GetRanking)
}
