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
	CreateThread(threadDetails []entities.ThreadDetail, title string, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error)
	GetThreads(token string) ([]entities.Thread, error)
	GetThread(id uint, token string) (entities.Thread, error)
	DeleteThread(thread_id uint) error
	// AddBookmark(thread_id uint, token string) (entities.Bookmark, error)
}

type threadService struct {
	repo         repositories.ThreadRepository
	bookmarkRepo repositories.BookmarkRepository
	favoriteRepo repositories.FavoriteRepository
}

func NewThreadUseCase(repo repositories.ThreadRepository, bookmarkRepo repositories.BookmarkRepository, favoriteRepo repositories.FavoriteRepository) ThreadUseCase {
	return &threadService{repo, bookmarkRepo, favoriteRepo}
}

func (service *threadService) CreateThread(threadDetails []entities.ThreadDetail, title string, token string, file multipart.FileHeader, c *fiber.Ctx) (entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Thread{}, err
	}
	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Thread{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
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

	thread, err := service.repo.CreateThread(user_id, title, imageUrl)
	if err != nil {
		return entities.Thread{}, err
	}

	for _, threadDetail := range threadDetails {
		threadDetail.ThreadID = thread.ID
		_, err := service.repo.CreateThreadDetail(threadDetail)
		if err != nil {
			return entities.Thread{}, err
		}
	}

	result, err := service.GetThread(thread.ID, token)
	if err != nil {
		return entities.Thread{}, err
	}

	thread_details, err := service.repo.GetThreadDetails(result.ID)
	if err != nil {
		return entities.Thread{}, err
	}

	result.Threads = thread_details

	return result, nil
}

func (service *threadService) GetThreads(token string) ([]entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return []entities.Thread{}, err
	}

	result, err := service.repo.GetThreads()
	if err != nil {
		return []entities.Thread{}, err
	}

	for i, thread := range result {

		bookmark, err := service.bookmarkRepo.FindBookMark(thread.ID, user_id)

		if err != nil {
			result[i].Bookmark = false
		} else {
			result[i].Bookmark = bookmark.Status
		}

		favorite, err := service.favoriteRepo.FindFavoriteThread(thread.ID, user_id)

		if err != nil {
			result[i].Favorite = false
		} else {
			result[i].Favorite = favorite.Status
		}

		if result[i].UserID != user_id {
			result[i].Owner = false
		} else {
			result[i].Owner = true
		}

		thread_details, err := service.repo.GetThreadDetails(thread.ID)
		if err != nil {
			return []entities.Thread{}, err
		}

		result[i].Threads = thread_details
	}

	return result, nil
}

func (service *threadService) GetThread(id uint, token string) (entities.Thread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Thread{}, err
	}

	result, err := service.repo.GetThread(id)
	if err != nil {
		return entities.Thread{}, err
	}

	thread_details, err := service.repo.GetThreadDetails(result.ID)
	if err != nil {
		return entities.Thread{}, err
	}

	result.Threads = thread_details

	bookmark, err := service.bookmarkRepo.FindBookMark(result.ID, user_id)

	if err != nil {
		result.Bookmark = false
	} else {
		result.Bookmark = bookmark.Status
	}

	if result.UserID != user_id {
		result.Owner = false
	} else {
		result.Owner = true
	}

	favorite, err := service.favoriteRepo.FindFavoriteThread(result.ID, user_id)

	if err != nil {
		result.Favorite = false
	} else {
		result.Favorite = favorite.Status
	}

	return result, nil
}

func (service *threadService) DeleteThread(thread_id uint) error {
	err := service.repo.DeleteThread(thread_id)
	if err != nil {
		return err
	}

	return nil
}
