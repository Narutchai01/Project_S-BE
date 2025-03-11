package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skin"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SkinRouters(app fiber.Router, admin fiber.Router, db *gorm.DB) {

	skinRepo := adapters.NewGormSkinRepository(db)
	skinService := usecases.NewSkinUseCase(skinRepo)
	skinHandler := adapters.NewHttpSkinHandler(skinService)

	skinAdmin := admin.Group("/skin")
	skinAdmin.Use(middlewares.AuthorizationRequired())
	skinAdmin.Post("/", skinHandler.CreateSkin)
	skinAdmin.Delete("/:id", skinHandler.DeleteSkin)
	skinAdmin.Put("/:id", skinHandler.UpdateSkin)

	skinUser := app.Group("/skin")
	skinUser.Get("/", skinHandler.GetSkins)
	skinUser.Get("/:id", skinHandler.GetSkin)

}
