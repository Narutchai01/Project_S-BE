package usecases

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecases interface {
	Register(user entities.User, c *fiber.Ctx) (entities.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUsecases {
	return &userService{repo}
}

func (service *userService) Register(user entities.User, c *fiber.Ctx) (entities.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return user, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = string(hashedPassword)

	return service.repo.CreateUser(user)
}