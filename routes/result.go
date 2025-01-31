package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ResultRoutes(app fiber.Router, db *gorm.DB) {

	resultRepo := adapters.NewGormResultRepository(db)
	resultService := usecases.NewResultUsecase(resultRepo)
	resultHandler := adapters.NewHttpResultHandler(resultService)

	app.Post("/", resultHandler.CreateResult)
	app.Get("/", resultHandler.GetResults)
}