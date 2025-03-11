package usecases

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FacialUsecases interface {
	CreateFacial(facial entities.Facial, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Facial, error)
	GetFacials() ([]entities.Facial, error)
	GetFacial(id int) (entities.Facial, error)
	UpdateFacial(id int, facial entities.Facial, file *multipart.FileHeader, c *fiber.Ctx) (entities.Facial, error)
	DeleteFacial(id int) error
}

type facialService struct {
	repo repositories.FacialRepository
}

func NewFacialUseCase(repo repositories.FacialRepository) FacialUsecases {
	return &facialService{repo}
}

func (service *facialService) CreateFacial(facial entities.Facial, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Facial, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Facial{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Facial{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/facial")

	if err != nil {
		return facial, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = os.Remove("./uploads/" + fileName)

	if err != nil {
		return facial, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	create_by, err := utils.ExtractToken(token)

	if err != nil {
		return facial, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	facial.CreateBY = create_by

	facial.Image = imageUrl

	return service.repo.CreateFacial(facial)
}

func (service *facialService) GetFacials() ([]entities.Facial, error) {
	return service.repo.GetFacials()
}

func (service *facialService) GetFacial(id int) (entities.Facial, error) {
	return service.repo.GetFacial(id)
}

func (service *facialService) UpdateFacial(id int, facial entities.Facial, file *multipart.FileHeader, c *fiber.Ctx) (entities.Facial, error) {
	oldValue, err := service.repo.GetFacial(id)
	if err != nil {
		return entities.Facial{}, errors.New("facial not found")
	}

	if file != nil {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.Facial{}, fmt.Errorf("failed to check directory: %w", err)
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.Facial{}, fmt.Errorf("failed to save file: %w", err)
		}

		imageUrl, err := utils.UploadImage(fileName, "/facial")

		if err != nil {
			return facial, err
		}

		err = os.Remove("./uploads/" + fileName)

		if err != nil {
			return facial, err
		}

		facial.Image = imageUrl
	}

	facial.ID = oldValue.ID
	facial.CreateBY = oldValue.CreateBY
	facial.Name = utils.CheckEmptyValueBeforeUpdate(facial.Name, oldValue.Name)
	facial.Image = utils.CheckEmptyValueBeforeUpdate(facial.Image, oldValue.Image)

	return service.repo.UpdateFacial(id, facial)

}

func (service *facialService) DeleteFacial(id int) error {

	_, err := service.repo.GetFacial(id)
	if err != nil {
		return errors.New("facial not found")
	}

	return service.repo.DeleteFacial(id)
}
