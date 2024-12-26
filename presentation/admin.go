package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/gofiber/fiber/v2"
)

type Admin struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

func ToAdminResponse(data entities.Admin) *fiber.Map {
	admin := Admin{
		ID:       data.ID,
		FullName: data.FullName,
		Email:    data.Email,
		Image:    data.Image,
	}

	return &fiber.Map{
		"status": true,
		"admin":  admin,
		"error":  nil,
	}
}

func ToAdminsResponse(data []entities.Admin) *fiber.Map {
	admins := []Admin{}

	for _, admin := range data {
		admins = append(admins, Admin{
			ID:       admin.ID,
			FullName: admin.FullName,
			Email:    admin.Email,
			Image:    admin.Image,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   admins,
		"error":  nil,
	}
}

func AdminErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"admin":  nil,
		"error":  err.Error(),
	}
}

func DeleteAdminResponse(id int) *fiber.Map {
	return &fiber.Map{
		"status":    true,
		"delete_id": id,
		"error":     nil,
	}
}

func AdminLoginResponse(token string, err error) *fiber.Map {
	if err != nil {
		return &fiber.Map{
			"status": false,
			"token":  nil,
			"error":  err.Error(),
		}
	}
	return &fiber.Map{
		"status": true,
		"token":  token,
		"error":  nil,
	}
}
