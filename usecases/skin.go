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

type SkinUsecases interface {
	CreateSkin(skin entities.Skin, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Skin, error)
	GetSkins() ([]entities.Skin, error)
	GetSkin(id int) (entities.Skin, error)
	UpdateSkin(id int, facial entities.Skin) (entities.Skin, error)
	DeleteSkin(id int) error
}

type skinService struct {
	repo repositories.SkinRepository
}

func NewSkinUseCase(repo repositories.SkinRepository) SkinUsecases {
	return &skinService{repo}
}

func (service *skinService) CreateSkin(skin entities.Skin, file multipart.FileHeader, c *fiber.Ctx, token string) (entities.Skin, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Skin{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Skin{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/skin")

	if err != nil {
		return skin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = os.Remove("./uploads/" + fileName)

	if err != nil {
		return skin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	create_by, err := utils.ExtractToken(token)

	if err != nil {
		return skin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	skin.CreateBY = create_by

	skin.Image = imageUrl

	return service.repo.CreateSkin(skin)
}

func (service *skinService) GetSkins() ([]entities.Skin, error) {
	return service.repo.GetSkins()
}

func (service *skinService) GetSkin(id int) (entities.Skin, error) {
	return service.repo.GetSkin(id)
}

func (service *skinService) UpdateSkin(id int, skin entities.Skin) (entities.Skin, error) {

	oldvalue, err := service.repo.GetSkin(id)

	if err != nil {
		return entities.Skin{}, err
	}
	skin.ID = oldvalue.ID
	skin.Name = utils.CheckEmptyValueBeforeUpdate(skin.Name, oldvalue.Name)
	skin.Image = utils.CheckEmptyValueBeforeUpdate(skin.Image, oldvalue.Image)
	return service.repo.UpdateSkin(id, skin)
}

func (service *skinService) UpdateSkinWithImage(id int, skin entities.Skin, file multipart.FileHeader, c *fiber.Ctx) (entities.Skin, error) {

	oldvalue, err := service.repo.GetSkin(id)
	if err != nil {
		return entities.Skin{}, err
	}

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Skin{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Skin{}, err
	}

	if oldvalue.Image == "" {
		imageUrl, err := utils.UploadImage(fileName, "/acne")
		if err != nil {
			return entities.Skin{}, fmt.Errorf("failed to upload new image: %w", err)
		}
		skin.Image = imageUrl
	} else {
		oldImage := path.Base(oldvalue.Image)
		err := utils.UpdateImage(oldImage, fileName)
		if err != nil {
			return entities.Skin{}, fmt.Errorf("failed to update existing image: %w", err)
		}

		skin.Image = oldvalue.Image
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return entities.Skin{}, fmt.Errorf("failed to remove temporary file: %w", err)
	}

	skin.ID = oldvalue.ID
	skin.Name = utils.CheckEmptyValueBeforeUpdate(skin.Name, oldvalue.Name)

	return service.repo.UpdateSkin(id, skin)
}

func (service *skinService) DeleteSkin(id int) error {
	return service.repo.DeleteSkin(id)
}
