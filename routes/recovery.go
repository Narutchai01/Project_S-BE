package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/recovery"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RecoveryRoutes(app fiber.Router, db *gorm.DB) {

	recoveryRepo := adapters.NewGormRecoveryRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	recoveryService := usecases.NewRecoveryUseCase(recoveryRepo, userRepo)
	recoveryHandler := adapters.NewHttpRecoveryHandler(recoveryService)

	app.Post("/", recoveryHandler.CreateRecovery)
	app.Post("/validation", recoveryHandler.ValidateRecovery)
	app.Post("/reset-password", recoveryHandler.ResetPassword)
}
