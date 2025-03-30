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
		Follower:      data.Follower,
		Following:     data.Following,
		SensitiveSkin: data.SensitiveSkin,
		Image:         data.Image,
		Follow:        data.Follow,
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

func PublicFollower(data entities.Follower) Follower {
	follow := Follower{
		ID:         uint64(data.ID),
		FollowerID: data.FollowerID,
		Follower:   *PublicUser(data.Follower),
		UserID:     data.UserID,
		User:       *PublicUser(data.User),
	}

	return follow
}

func ToFollowerResponse(data entities.Follower) *Responses {
	follow := PublicFollower(data)

	return &Responses{
		Status: true,
		Data:   follow,
		Error:  nil,
	}
}
