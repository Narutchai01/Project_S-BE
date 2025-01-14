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
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecases interface {
	CreateAdmin(admin entities.Admin, file multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error)
	GetAdmins() ([]entities.Admin, error)
	GetAdmin(id int) (entities.Admin, error)
	UpdateAdmin(token string, admin entities.Admin, file *multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error)
	DeleteAdmin(id int) (entities.Admin, error)
	LogIn(email string, password string) (string, error)
	GetAdminByToken(token string) (entities.Admin, error)
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminUseCase(repo repositories.AdminRepository) AdminUsecases {
	return &adminService{repo}
}

func (service *adminService) CreateAdmin(admin entities.Admin, file multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := utils.CheckDirectoryExist(); err != nil {
		return entities.Admin{}, err
	}

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Admin{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/admin")

	if err != nil {
		return admin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return admin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		return admin, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	admin.Password = string(hashedPassword)

	admin.Image = imageUrl

	return service.repo.CreateAdmin(admin)
}

func (service *adminService) GetAdmins() ([]entities.Admin, error) {
	return service.repo.GetAdmins()
}

func (service *adminService) GetAdmin(id int) (entities.Admin, error) {
	return service.repo.GetAdmin(id)
}

func (service *adminService) UpdateAdmin(token string, admin entities.Admin, file *multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error) {

	id, err := utils.ExtractToken(token)
	if err != nil {
		return admin, fmt.Errorf("failed to extract token: %w", err)
	}

	oldamin, err := service.repo.GetAdmin(int(id))
	if err != nil {
		return entities.Admin{}, err
	}

	if file != nil {
		fileName := uuid.New().String() + ".jpg"

		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.Admin{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.Admin{}, err
		}

		if oldamin.Image == "" {
			imageUrl, err := utils.UploadImage(fileName, "/admin")
			if err != nil {
				return entities.Admin{}, fmt.Errorf("failed to upload new image: %w", err)
			}
			admin.Image = imageUrl
		} else {
			oldImage := path.Base(oldamin.Image)
			fmt.Println("Old Image Path:", oldImage)
			err := utils.UpdateImage(oldImage, fileName, "admin")
			if err != nil {
				return entities.Admin{}, fmt.Errorf("failed to update existing image: %w", err)
			}

			admin.Image = oldamin.Image
		}

		err = os.Remove("./uploads/" + fileName)
		if err != nil {
			return entities.Admin{}, fmt.Errorf("failed to remove temporary file: %w", err)
		}
	}

	admin.ID = oldamin.ID

	admin.Image = utils.CheckEmptyValueBeforeUpdate(admin.Image, oldamin.Image)
	admin.Password = utils.CheckEmptyValueBeforeUpdate(admin.Password, oldamin.Password)
	admin.FullName = utils.CheckEmptyValueBeforeUpdate(admin.FullName, oldamin.FullName)
	admin.Email = utils.CheckEmptyValueBeforeUpdate(admin.Email, oldamin.Email)

	return service.repo.UpdateAdmin(int(id), admin)
}

func (service *adminService) DeleteAdmin(id int) (entities.Admin, error) {

      old_admin, err := service.repo.GetAdmin(id)
      if err != nil {
            return entities.Admin{}, err
      }

      oldImage := path.Base(old_admin.Image)
      if err := utils.DeleteImage(oldImage, "admin"); err != nil {
            return entities.Admin{}, fmt.Errorf("failed to update existing image: %w", err)
      }

      return service.repo.DeleteAdmin(id)
}

func (service *adminService) LogIn(email string, password string) (string, error) {
	admin, err := service.repo.GetAdminByEmail(email)

	if err != nil {
		return "something wrong!", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return "something wrong!", err
	}

	// create jwt token
	token, err := utils.GenerateToken(int(admin.ID))

	if err != nil {
		return "something wrong!", err
	}

	return token, nil
}

func (service *adminService) GetAdminByToken(token string) (entities.Admin, error) {
	id, err := utils.ExtractToken(token)

	if err != nil {
		return entities.Admin{}, err
	}

	return service.repo.GetAdmin(int(id))
}