package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type UserRepository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	UpdateUserPasswordById(id int, newPassword string) (entities.User, error)
	GetUser(id uint) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	Follower(follow_id uint, user_id uint) (entities.Follower, error)
	FindFollower(follow_id uint, user_id uint) (entities.Follower, error)
	DeleteFollower(id uint) (entities.Follower, error)
}
