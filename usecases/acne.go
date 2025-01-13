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

type AcneUseCase interface {
	CreateAcne(acne entities.Acne, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Acne, error)
	GetAcnes() ([]entities.Acne, error)
	GetAcne(id int) (entities.Acne, error)
}

type acneService struct {
	repo repositories.AcneRepository
}

func NewAcneUseCase(repo repositories.AcneRepository) AcneUseCase {
	return &acneService{repo}
}

func (service *acneService) CreateAcne(acne entities.Acne, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Acne, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Acne{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Acne{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/acne")

	if err != nil {
		return acne, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	err = os.Remove("./uploads/" + fileName)

	if err != nil {
		return acne, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	create_by, err := utils.ExtractToken(token)

	if err != nil {
		return acne, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	acne.Image = imageUrl

	acne.CreateBY = create_by

	return service.repo.CreateAcne(acne)
}

func (service *acneService) GetAcnes() ([]entities.Acne, error) {
	return service.repo.GetAcnes()
}

func (service *acneService) GetAcne(id int) (entities.Acne, error) {
	return service.repo.GetAcne(id)
}
