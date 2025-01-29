package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/acne"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AcneRouters(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	acneRepo := adapters.NewGormAcneRepository(db)
	acneService := usecases.NewAcneUseCase(acneRepo)
	acneHandler := adapters.NewHttpAcneHandler(acneService)

	acneUser := app.Group("/acne")
	acneUser.Get("/", acneHandler.GetAcnes)
	acneUser.Get("/:id", acneHandler.GetAcne)
	
	acneAdmin := admin.Group("/acne")
	acneAdmin.Use(middlewares.AuthorizationRequired())
	acneAdmin.Post("/", acneHandler.CreateAcne)
	acneAdmin.Delete("/:id", acneHandler.DeleteAcne)
	acneAdmin.Put("/:id", acneHandler.UpdateAcne)
}
