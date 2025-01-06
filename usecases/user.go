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

type UserUsecases interface {
	Register(user entities.User, file multipart.FileHeader, c *fiber.Ctx) (entities.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUsecases {
	return &userService{repo}
}

func (service *userService) Register(user entities.User, file multipart.FileHeader, c *fiber.Ctx) (entities.User, error) {

	fileName := uuid.New().String() + ".jpg"

	//Aut add this
	dir := "./uploads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return entities.User{}, err
		}
	}
	//

	if err := c.SaveFile(&file, "./uploads/"+fileName); err != nil {
		return entities.User{}, err
	}

	imageUrl, err := utils.UploadImage(fileName, "/")

	if err != nil {
		return user, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = os.Remove("./uploads/" + fileName)
	if err != nil {
		return user, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return user, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	user.Image = imageUrl

	return service.repo.CreateUser(user)
}