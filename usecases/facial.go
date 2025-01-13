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

type FacialUsecases interface {
	CreateFacial(facial entities.Facial, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Facial, error)
	GetFacials() ([]entities.Facial, error)
	GetFacial(id int) (entities.Facial, error)
	UpdateFacial(id int, facial entities.Facial) (entities.Facial, error)
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

func (service *facialService) UpdateFacial(id int, facial entities.Facial) (entities.Facial, error) {
	return service.repo.UpdateFacial(id, facial)
}

func (service *facialService) DeleteFacial(id int) error {
	return service.repo.DeleteFacial(id)
}
