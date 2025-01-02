package usecases

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SkincareUsecases interface {
	CreateSkincare(skincare entities.Skincare, file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Skincare, error)
	GetSkincares() ([]entities.Skincare, error)
	GetSkincare(id int) (entities.Skincare, error)
	UpdateSkincare(id int, skincare entities.Skincare) (entities.Skincare, error)
	DeleteSkincare(id int) (entities.Skincare, error)
}

type skincareService struct {
	repo repositories.SkincareRepository
}

func NewSkincareUseCase(repo repositories.SkincareRepository) SkincareUsecases {
	return &skincareService{repo}
}

func (service *skincareService) CreateSkincare(skincare entities.Skincare, file multipart.FileHeader, token string, c *fiber.Ctx) (entities.Skincare, error) {

	fileName := uuid.New().String() + ".jpg"

	//Aut add this
	dir := "./uploads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		    return entities.Skincare{}, err
		}
	}
	//

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Skincare{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/")

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

	log.Println("Extracted user_id from token: ", create_by_id)
	skincare.CreateBY = create_by_id

	return service.repo.CreateSkincare(skincare)
}

func (service *skincareService) GetSkincares() ([]entities.Skincare, error) {
	return service.repo.GetSkincares()
}

func (service *skincareService) GetSkincare(id int) (entities.Skincare, error) {
	return service.repo.GetSkincare(id)
}

func (service *skincareService) UpdateSkincare(id int, skincare entities.Skincare) (entities.Skincare, error) {

	old_skincare, err := service.repo.GetSkincare(id)

	skincare.ID = old_skincare.ID

	if err != nil {
		return entities.Skincare{}, err
	}

	skincare.Name = utils.CheckEmptyValueBeforeUpdate(skincare.Name, old_skincare.Name)
	skincare.Image = utils.CheckEmptyValueBeforeUpdate(skincare.Image, old_skincare.Image)
	skincare.Description = utils.CheckEmptyValueBeforeUpdate(skincare.Description, old_skincare.Description)

	return service.repo.UpdateSkincare(id, skincare)
}

func (service *skincareService) DeleteSkincare(id int) (entities.Skincare, error) {
	return service.repo.DeleteSkincare(id)
}
