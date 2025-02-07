package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func resultRoutes(app fiber.Router, db *gorm.DB) {
	resultRepo := adapters.NewGormResultRepository(db)
	resultService := usecases.NewResultsUsecase(resultRepo)
	resultHandler := adapters.NewHttpResultHandler(resultService)

	resultUser := app.Group("/result")
	resultUser.Post("/", resultHandler.CreateResult)
	resultUser.Get("/", resultHandler.GetResults)
	resultUser.Get("/:id", resultHandler.GetResult)

}
