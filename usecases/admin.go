package usecases

import (
	"mime/multipart"
	"os"

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
	UpdateAdmin(id int, admin entities.Admin) (entities.Admin, error)
	DeleteAdmin(id int) (entities.Admin, error)
	LogIn(email string, password string) (entities.Admin, error)
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminUseCase(repo repositories.AdminRepository) AdminUsecases {
	return &adminService{repo}
}

func (service *adminService) CreateAdmin(admin entities.Admin, file multipart.FileHeader, c *fiber.Ctx) (entities.Admin, error) {

	fileName := uuid.New().String() + ".jpg"

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.Admin{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/")

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

func (service *adminService) UpdateAdmin(id int, admin entities.Admin) (entities.Admin, error) {

	oldamin, err := service.repo.GetAdmin(id)

	admin.ID = oldamin.ID

	if err != nil {
		return admin, err
	}

	if admin.Image == "" {
		admin.Image = oldamin.Image
	}

	if admin.Password == "" {
		admin.Password = oldamin.Password
	}

	if admin.FullName == "" {
		admin.FullName = oldamin.FullName
	}

	return service.repo.UpdateAdmin(id, admin)
}

func (service *adminService) DeleteAdmin(id int) (entities.Admin, error) {
	return service.repo.DeleteAdmin(id)
}

func (service *adminService) LogIn(email string, password string) (entities.Admin, error) {
	admin, err := service.repo.GetAdminByEmail(email)

	if err != nil {
		return admin, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return admin, err
	}

	return admin, nil
}
