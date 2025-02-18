package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ThreadRouters(app fiber.Router, db *gorm.DB) {

	threadRepo := adapters.NewGormThreadRepository(db)
	threadService := usecases.NewThreadUseCase(threadRepo)
	threadHandler := adapters.NewHttpThreadHandler(threadService)

	threadGroup := app.Group("/thread")

	threadGroup.Post("/", threadHandler.CreateThread)

}
