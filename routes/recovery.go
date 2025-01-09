package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/recovery"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RecoveryRoutes(app fiber.Router, db *gorm.DB) {

	recoveryRepo := adapters.NewGormRecoveryRepository(db)
	recoveryService := usecases.NewRecoveryUseCase(recoveryRepo)
	recoveryHandler := adapters.NewHttpRecoveryHandler(recoveryService)

	app.Post("/", recoveryHandler.CreateRecovery)
	app.Delete("/:id", recoveryHandler.DeleteRecoveryById)
	app.Get("/", recoveryHandler.GetRecoveries)
	app.Post("/validation", recoveryHandler.OtpValidation)
}