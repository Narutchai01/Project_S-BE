package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SkincareRoutes(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	skincareRepo := adapters.NewGormSkincareRepository(db)
	skincareService := usecases.NewSkincareUseCase(skincareRepo)
	skincareHandler := adapters.NewHttpSkincareHandler(skincareService)

	//user

	userSkincare := app.Group("/skincare")
	userSkincare.Get("/", skincareHandler.GetSkincares)
	userSkincare.Get("/:id", skincareHandler.GetSkincareById)

	//admin
	adminSkincare := admin.Group("/skincare")
	adminSkincare.Use(middlewares.AuthorizationRequired())
	adminSkincare.Post("/", skincareHandler.CreateSkincare)
	adminSkincare.Put("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	adminSkincare.Put("/:id", skincareHandler.UpdateSkincareById)
	adminSkincare.Delete("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	adminSkincare.Delete("/:id", skincareHandler.DeleteSkincareById)
}
