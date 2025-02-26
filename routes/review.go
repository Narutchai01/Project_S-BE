package routes

import (
	adaptersReview "github.com/Narutchai01/Project_S-BE/adapters/review"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ReviewRoutes(app fiber.Router, db *gorm.DB) {
	reviewRepo := adaptersReview.NewGormReviewRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	reviewService := usecases.NewReviewThreadUseCase(reviewRepo, userRepo)
	reviewHandler := adaptersReview.NewHttpReviewRepository(reviewService)

	reviewGroup := app.Group("/reviews").Use(middlewares.AuthorizationRequired())
	reviewGroup.Post("/", reviewHandler.CreateReviewSkincare)

}
