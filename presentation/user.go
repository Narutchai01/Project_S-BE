package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
)

func UserResponse(data entities.User) *fiber.Map {
	user := User{
		ID:            data.ID,
		FullName:      data.FullName,
		Email:         data.Email,
		Birthday:      data.Birthday,
		SensitiveSkin: data.SensitiveSkin,
		Image:         data.Image,
	}

	return &fiber.Map{
		"status": true,
		"user":   user,
		"error":  nil,
	}
}

func MiniProfileUserResponse(data entities.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		FullName: data.FullName,
		Image:    data.Image,
	}

	return &fiber.Map{
		"status": true,
		"user":   user,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"user":   nil,
		"error":  err.Error(),
	}
}

func UserLoginResponse(token string, err error) *Responses {
	if err != nil {
		return &Responses{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		}
	}
	return &Responses{
		Status: true,
		Data: map[string]string{
			"token": token,
		},
		Error: nil,
	}
}
