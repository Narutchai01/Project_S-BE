package presentation

import (
	"github.com/Narutchai01/Project_S-BE/entities"
)

func PublicUser(data entities.User) *User {
	user := User{
		ID:            data.ID,
		FullName:      data.FullName,
		Email:         data.Email,
		Birthday:      data.Birthday,
		SensitiveSkin: data.SensitiveSkin,
		Image:         data.Image,
	}

	return &user
}

func UserResponse(data entities.User) *Responses {

	user := PublicUser(data)

	return &Responses{
		Status: true,
		Data:   user,
		Error:  nil,
	}
}

func MiniProfileUserResponse(data entities.User) *Responses {
	user := User{
		ID:       data.ID,
		FullName: data.FullName,
		Image:    data.Image,
	}

	return &Responses{
		Status: true,
		Data:   user,
		Error:  nil,
	}
}
