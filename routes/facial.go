package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/facial"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FacialRouters(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	facialRepo := adapters.NewGormFacialRepository(db)
	facialService := usecases.NewFacialUseCase(facialRepo)
	facialHandler := adapters.NewHttpFacialHandler(facialService)

	facialAdmin := admin.Group("/facial")
	facialAdmin.Use(middlewares.AuthorizationRequired())
	facialAdmin.Post("/", facialHandler.CreateFacial)
	facialAdmin.Delete("/:id", facialHandler.DeleteFacial)
	facialAdmin.Put("/:id", facialHandler.UpdateFacial)

	facialUser := app.Group("/facial")
	facialUser.Get("/", facialHandler.GetFacials)
	facialUser.Get("/:id", facialHandler.GetFacial)

}
