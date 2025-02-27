package usecases

import (
	"mime/multipart"
	"os"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ThreadUseCase interface {
	CreateThread(thread entities.Thread, token string, files []*multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error)
	GetThread(thread_id uint, token string) (entities.Thread, error)
	GetThreads(token string) ([]entities.Thread, error)
}

type threadService struct {
	threadRepo   repositories.ThreadRepository
	userRepo     repositories.UserRepository
	favoriteRepo repositories.FavoriteRepository
	bookmarkRepo repositories.BookmarkRepository
}

func NewThreadUseCase(threadRepo repositories.ThreadRepository, userRepo repositories.UserRepository, favoriteRepo repositories.FavoriteRepository, bookmarkRepo repositories.BookmarkRepository) ThreadUseCase {
	return &threadService{threadRepo, userRepo, favoriteRepo, bookmarkRepo}
}

func (service *threadService) CreateThread(thread entities.Thread, token string, files []*multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Thread{}, err
	}

	_, err = service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Thread{}, err
	}

	thread.UserID = user_id

	thread, err = service.threadRepo.CreateThread(thread)

	if err != nil {
		return entities.Thread{}, err
	}

	ImageURLs := []string{}

	for _, file := range files {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.Thread{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.Thread{}, err
		}

		imageUrl, err := utils.UploadImage(fileName, "/thread")

		if err != nil {
			return entities.Thread{}, err
		}

		err = os.Remove("./uploads/" + fileName)

		if err != nil {
			return entities.Thread{}, err
		}

		ImageURLs = append(ImageURLs, imageUrl)
	}

	for _, imageUrl := range ImageURLs {
		image := entities.ThreadImage{
			ThreadID: thread.ID,
			Image:    imageUrl,
		}

		_, err := service.threadRepo.CreateThreadImage(image)

		if err != nil {
			return entities.Thread{}, err
		}
	}

	thread, err = service.threadRepo.GetThread(thread.ID)

	if err != nil {
		return entities.Thread{}, err
	}

	Images, err := service.threadRepo.GetThreadImages(thread.ID)

	if err != nil {
		return entities.Thread{}, err
	}

	thread.Images = Images

	favorite, err := service.favoriteRepo.FindFavoriteThread(user_id, thread.ID)

	if err != nil {
		thread.Favorite = false
	} else {
		thread.Favorite = favorite.Status
	}

	favoriteCount, err := service.favoriteRepo.CountFavoriteThread(thread.ID)

	if err != nil {
		thread.FavoriteCount = 0
	}

	thread.FavoriteCount = favoriteCount

	bookmark, err := service.bookmarkRepo.FindBookMarkThread(thread.ID, user_id)

	if err != nil {
		thread.Bookmark = false
	} else {
		thread.Bookmark = bookmark.Status
	}

	if user_id != thread.UserID {
		thread.Owner = false
	} else {
		thread.Owner = true
	}

	return thread, nil
}

func (service *threadService) GetThread(thread_id uint, token string) (entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.Thread{}, err
	}

	thread, err := service.threadRepo.GetThread(thread_id)
	if err != nil {
		return entities.Thread{}, fiber.NewError(fiber.StatusNotFound, "thread not found")
	}

	thread_images, err := service.threadRepo.GetThreadImages(thread.ID)

	if err != nil {
		return entities.Thread{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	thread.Images = thread_images

	favorite, err := service.favoriteRepo.FindFavoriteThread(thread_id, user_id)
	if err != nil {
		thread.Favorite = false
	} else {
		thread.Favorite = favorite.Status
	}

	favoriteCount, err := service.favoriteRepo.CountFavoriteThread(thread_id)
	if err != nil {
		thread.FavoriteCount = 0
	} else {
		thread.FavoriteCount = favoriteCount
	}

	bookmark, err := service.bookmarkRepo.FindBookMarkThread(thread_id, user_id)
	if err != nil {
		thread.Bookmark = false
	} else {
		thread.Bookmark = bookmark.Status
	}

	if user_id != thread.UserID {
		thread.Owner = false
	} else {
		thread.Owner = true
	}

	return thread, nil
}

func (service *threadService) GetThreads(token string) ([]entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Thread{}, err
	}

	threads, err := service.threadRepo.GetThreads()
	if err != nil {
		return []entities.Thread{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	for i, thread := range threads {
		thread_images, err := service.threadRepo.GetThreadImages(thread.ID)

		if err != nil {
			return []entities.Thread{}, fiber.NewError(fiber.StatusInternalServerError, "internal server error")
		}

		threads[i].Images = thread_images

		favorite, err := service.favoriteRepo.FindFavoriteThread(thread.ID, user_id)

		if err != nil {
			threads[i].Favorite = false
		} else {
			threads[i].Favorite = favorite.Status
		}

		favoriteCount, err := service.favoriteRepo.CountFavoriteThread(thread.ID)

		if err != nil {
			threads[i].FavoriteCount = 0
		} else {
			threads[i].FavoriteCount = favoriteCount
		}

		bookmark, err := service.bookmarkRepo.FindBookMarkThread(thread.ID, user_id)
		if err != nil {
			threads[i].Bookmark = false
		} else {
			threads[i].Bookmark = bookmark.Status
		}

		if user_id != thread.UserID {
			threads[i].Owner = false
		} else {
			threads[i].Owner = true
		}
	}

	return threads, nil
}
