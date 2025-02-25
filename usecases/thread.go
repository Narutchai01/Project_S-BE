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
}

type threadService struct {
	threadRepo repositories.ThreadRepository
	userRepo   repositories.UserRepository
}

func NewThreadUseCase(threadRepo repositories.ThreadRepository, userRepo repositories.UserRepository) ThreadUseCase {
	return &threadService{threadRepo, userRepo}
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

	return thread, nil
}
