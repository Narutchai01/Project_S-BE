package routes

import (
	adaptersBookmark "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adaptersFav "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ThreadRouters(app fiber.Router, db *gorm.DB) {

	threadRepo := adapters.NewGormThreadRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	favoriteRepo := adaptersFav.NewGormFavoriteRepository(db)
	bookmarkRepo := adaptersBookmark.NewGormBookmarkRepository(db)
	threadService := usecases.NewThreadUseCase(threadRepo, userRepo, favoriteRepo, bookmarkRepo)
	communityService := usecases.NewCommunityUseCase(communityRepo, userRepo, favoriteRepo, bookmarkRepo)
	threadHandler := adapters.NewHttpThreadRepository(threadService, communityService)

	threadGroup := app.Group("/thread")
	threadGroup.Post("/", threadHandler.CreateThread)
	threadGroup.Get("/", threadHandler.GetThreads)
	threadGroup.Get("/:id", threadHandler.GetThread)

}
