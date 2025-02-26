package usecases

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SkincareUsecases interface {
	CreateSkincare(skincare entities.Skincare, file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Skincare, error)
	GetSkincares() ([]entities.Skincare, error)
	GetSkincareById(id int) (entities.Skincare, error)
	UpdateSkincareById(id int, skincare entities.Skincare, file *multipart.FileHeader, c *fiber.Ctx) (entities.Skincare, error)
	DeleteSkincareById(id int) (entities.Skincare, error)
}

type skincareService struct {
	repo repositories.SkincareRepository
}

func NewSkincareUseCase(repo repositories.SkincareRepository) SkincareUsecases {
	return &skincareService{repo}
}

func (service *skincareService) CreateSkincare(skincare entities.Skincare, file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Skincare, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Skincare{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Skincare{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/skincare")

	if err != nil {
		return skincare, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return skincare, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	skincare.Image = imageUrl

	// ดึง user_id จาก Token
	create_by_id, err := utils.ExtractToken(token)
	if err != nil {
		return skincare, fmt.Errorf("failed to extract token: %w", err)
	}

	skincare.CreateBY = create_by_id

	return service.repo.CreateSkincare(skincare)
}

func (service *skincareService) GetSkincares() ([]entities.Skincare, error) {
	return service.repo.GetSkincares()
}

func (service *skincareService) GetSkincareById(id int) (entities.Skincare, error) {
	return service.repo.GetSkincareById(id)
}

func (service *skincareService) UpdateSkincareById(id int, skincare entities.Skincare, file *multipart.FileHeader, c *fiber.Ctx) (entities.Skincare, error) {

	old_skincare, err := service.repo.GetSkincareById(id)
	if err != nil {
		return entities.Skincare{}, err
	}

	if file != nil {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.Skincare{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.Skincare{}, err
		}

		imageUrl, err := utils.UploadImage(fileName, "/skincare")

		if err != nil {
			return skincare, err
		}

		err = os.Remove("./uploads/" + fileName)

		if err != nil {
			return skincare, err
		}

		skincare.Image = imageUrl
	}

	skincare.ID = old_skincare.ID

	skincare.Name = utils.CheckEmptyValueBeforeUpdate(skincare.Name, old_skincare.Name)
	skincare.Image = utils.CheckEmptyValueBeforeUpdate(skincare.Image, old_skincare.Image)
	skincare.Description = utils.CheckEmptyValueBeforeUpdate(skincare.Description, old_skincare.Description)

	return service.repo.UpdateSkincareById(id, skincare)
}

func (service *skincareService) DeleteSkincareById(id int) (entities.Skincare, error) {

	old_skincare, err := service.repo.GetSkincareById(id)
	if err != nil {
		return entities.Skincare{}, err
	}

	oldImage := path.Base(old_skincare.Image)
	if err := utils.DeleteImage(oldImage, "skincare"); err != nil {
		return entities.Skincare{}, fmt.Errorf("failed to update existing image: %w", err)
	}

	return service.repo.DeleteSkincareById(id)
}
