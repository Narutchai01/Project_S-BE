package usecases

import (
	"fmt"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecases interface {
	Register(user entities.User, c *fiber.Ctx) (entities.User, error)
	LogIn(email string, password string) (string, error)
	ChangePassword(id int, ewPassword string, c *fiber.Ctx) (entities.User, error)
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

func (service *userService) LogIn(email string, password string) (string, error) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return "something wrong!", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "something wrong!", err
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		return "something wrong!", err
	}

	return token, nil
}

func (service *userService) ChangePassword(id int, newPassword string, c *fiber.Ctx) (entities.User, error) {

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to hashed password: %w", err)
	}

	return service.repo.UpdateUserPasswordById(id, string(hashedNewPassword))
}