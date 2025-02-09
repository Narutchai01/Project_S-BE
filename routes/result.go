package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ResultRoutes(app fiber.Router, db *gorm.DB) {
	resultRepo := adapters.NewGormResultRepository(db)
	resultService := usecases.NewResultsUsecase(resultRepo)
	resultHandler := adapters.NewHttpResultHandler(resultService)

	resultGroup := app.Group("/results")
	resultGroup.Post("/", resultHandler.CreateResult)
	resultGroup.Get("/", resultHandler.GetResults)
	resultGroup.Get("/latest", resultHandler.GetResultLatest)
	resultGroup.Get("/:id", resultHandler.GetResult)
	resultGroup.Put("/:id", resultHandler.UpdateResult)
	resultGroup.Delete("/:id", resultHandler.DeleteResult)
}
