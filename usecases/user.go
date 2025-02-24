package usecases

import (
	"fmt"
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
	Register(user entities.User, c *fiber.Ctx) (entities.User, error)
	LogIn(email string, password string) (string, error)
	ChangePassword(id int, ewPassword string, c *fiber.Ctx) (entities.User, error)
	GoogleSignIn(user entities.User) (string, error)
	GetUser(token string) (entities.User, error)
	UpdateUser(user entities.User, token string, file *multipart.FileHeader, c *fiber.Ctx) (entities.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) UserUsecases {
	return &userService{repo}
}

func (service *userService) Register(user entities.User, c *fiber.Ctx) (entities.User, error) {

	_, err := service.repo.GetUserByEmail(user.Email)
	if err == nil {
		return user, fmt.Errorf("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, fmt.Errorf("failed to hashed password: %w", err)
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

func (service *userService) GoogleSignIn(user entities.User) (string, error) {
	email := user.Email
	existingUser, err := service.repo.GetUserByEmail(email)

	if err != nil {
		newUser, err := service.repo.CreateUser(user)
		if err != nil {
			return "", err
		}
		return utils.GenerateToken(int(newUser.ID))
	}

	return utils.GenerateToken(int(existingUser.ID))
}

func (service *userService) GetUser(token string) (entities.User, error) {
	id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.User{}, err
	}
	return service.repo.GetUser(uint(id))
}

func (service *userService) UpdateUser(user entities.User, token string, file *multipart.FileHeader, c *fiber.Ctx) (entities.User, error) {

	id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.User{}, err
	}
	user.ID = uint(id)

	old_value, err := service.repo.GetUser(uint(id))

	if err != nil {
		return entities.User{}, err
	}

	if file != nil {
		fileName := uuid.New().String() + ".jpg"
		if err := utils.CheckDirectoryExist(); err != nil {
			return entities.User{}, err
		}

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return entities.User{}, err
		}

		imageUrl, err := utils.UploadImage(fileName, "/user")

		if err != nil {
			return entities.User{}, err
		}

		err = os.Remove("./uploads/" + fileName)

		if err != nil {
			return entities.User{}, err
		}

		user.Image = imageUrl
	}

	user.ID = old_value.ID
	user.Email = utils.CheckEmptyValueBeforeUpdate(user.Email, old_value.Email)
	user.FullName = utils.CheckEmptyValueBeforeUpdate(user.FullName, old_value.FullName)
	user.Image = utils.CheckEmptyValueBeforeUpdate(user.Image, old_value.Image)
	user.Password = utils.CheckEmptyValueBeforeUpdate(user.Password, old_value.Password)

	return service.repo.UpdateUser(user)

}
