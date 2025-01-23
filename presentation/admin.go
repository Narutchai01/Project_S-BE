package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func ToAdminResponse(data entities.Admin) *Responses {
	admin := Admin{
		ID:       data.ID,
		FullName: data.FullName,
		Email:    data.Email,
		Image:    data.Image,
	}

	return &Responses{
		Status: true,
		Data:   admin,
		Error:  nil,
	}
}

func ToAdminsResponse(data []entities.Admin) *Responses {
	admins := []Admin{}

	for _, admin := range data {
		admins = append(admins, Admin{
			ID:       admin.ID,
			FullName: admin.FullName,
			Email:    admin.Email,
			Image:    admin.Image,
		})
	}
	return &Responses{
		Status: true,
		Data:   admins,
		Error:  nil,
	}
}
