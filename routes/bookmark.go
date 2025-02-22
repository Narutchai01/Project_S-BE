package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookMarkRouters(app fiber.Router, db *gorm.DB) {

	bookmarkRepo := adapters.NewGormBookmarkRepository(db)
	bookmarkService := usecases.NewBookmarkUseCase(bookmarkRepo)
	bookmarkHandler := adapters.NewHttpBookmarkHandler(bookmarkService)

	BookmarkGroup := app.Group("/bookmark").Use(middlewares.AuthorizationRequired())

	BookmarkGroup.Post("/:id", bookmarkHandler.BookMarkThread)
}
