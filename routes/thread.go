package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ThreadRouters(app fiber.Router, db *gorm.DB) {

	threadRepo := adapters.NewGormThreadRepository(db)
	threadService := usecases.NewThreadUseCase(threadRepo)
	threadHandler := adapters.NewHttpThreadHandler(threadService)

	threadGroup := app.Group("/thread")
	threadGroup.Post("/", middlewares.AuthorizationRequired(), threadHandler.CreateThread)
	threadGroup.Get("/", threadHandler.GetThreads)
	threadGroup.Get("/:id", threadHandler.GetThread)
	threadGroup.Post("/:id/bookmark", threadHandler.BookMark)
	threadGroup.Delete("/:id", threadHandler.DeleteThread)

}
