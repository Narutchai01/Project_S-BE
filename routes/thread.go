package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ThreadRouters(app fiber.Router, db *gorm.DB) {

	threadRepo := adapters.NewGormThreadRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	threadService := usecases.NewThreadUseCase(threadRepo, userRepo)
	threadHandler := adapters.NewHttpThreadRepository(threadService)

	threadGroup := app.Group("/thread")
	threadGroup.Post("/", threadHandler.CreateThread)

}
