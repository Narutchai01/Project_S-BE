package routes

import (
	adaptersBookmark "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersFavorite "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ThreadRouters(app fiber.Router, db *gorm.DB) {

	threadRepo := adapters.NewGormThreadRepository(db)
	bookmarkRepo := adaptersBookmark.NewGormBookmarkRepository(db)
	favoriteRepo := adaptersFavorite.NewGormFavoriteRepository(db)
	threadService := usecases.NewThreadUseCase(threadRepo, bookmarkRepo, favoriteRepo)
	threadHandler := adapters.NewHttpThreadHandler(threadService)

	threadGroup := app.Group("/thread")
	threadGroup.Post("/", middlewares.AuthorizationRequired(), threadHandler.CreateThread)
	threadGroup.Get("/", threadHandler.GetThreads)
	threadGroup.Get("/:id", threadHandler.GetThread)
	// threadGroup.Post("/:id/bookmark", threadHandler.BookMark)
	threadGroup.Delete("/:id", threadHandler.DeleteThread)
	threadGroup.Put("/:id", threadHandler.UpdateThread)

}
