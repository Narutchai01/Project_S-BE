package routes

import (
	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	adaptersCommunity "github.com/Narutchai01/Project_S-BE/adapters/community"
	adaptersFav "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	adaptersUser "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/middlewares"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookMarkRouters(app fiber.Router, db *gorm.DB) {

	bookmarkRepo := adapters.NewGormBookmarkRepository(db)
	userRepo := adaptersUser.NewGormUserRepository(db)
	communityRepo := adaptersCommunity.NewGormCommunityRepository(db)
	favoriteRepo := adaptersFav.NewGormFavoriteRepository(db)
	bookmarkService := usecases.NewBookmarkUseCase(bookmarkRepo, userRepo, communityRepo, favoriteRepo)
	bookmarkHandler := adapters.NewHttpBookmarkHandler(bookmarkService)

	BookmarkGroup := app.Group("/bookmark").Use(middlewares.AuthorizationRequired())

	BookmarkGroup.Post("/thread/:id", bookmarkHandler.BookMarkThread)
	BookmarkGroup.Post("/review/:id", bookmarkHandler.BookMarkReviewSkincare)
	BookmarkGroup.Get("/get/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(presentation.ErrorResponse(fiber.ErrNotFound))
	})
	BookmarkGroup.Get("/get/:user_id", bookmarkHandler.GetCommunitiesBookmark)
}
