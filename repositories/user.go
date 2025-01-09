package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type UserRepository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
}