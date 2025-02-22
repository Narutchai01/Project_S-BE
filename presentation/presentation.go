package presentation

import (
	"time"

	"github.com/Narutchai01/Project_S-BE/entities"
)

type Responses struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

type Admin struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
}

type Acne struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	CreateBY uint   `json:"create_by"`
}

type Recovery struct {
	ID     uint   `json:"id"`
	UserId int    `json:"user_id"`
	OTP    string `json:"otp"`
}

type Skincare struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreateBY    uint   `json:"create_by"`
}

type User struct {
	ID            uint       `json:"id"`
	FullName      string     `json:"full_name"`
	Email         string     `json:"email"`
	Birthday      *time.Time `json:"birthday"`
	SensitiveSkin *bool      `json:"sensitive_skin"`
	Image         string     `json:"image"`
}

type Facial struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	CreateBY uint   `json:"create_by"`
}
type Result struct {
	ID         uint                  `json:"id"`
	UserID     uint                  `json:"user_id"`
	SkincareID []uint                `json:"skincare_id"`
	SkinID     uint                  `json:"skin_id"`
	Skin       Skin                  `json:"skin"`
	Image      string                `json:"image"`
	Skincare   []Skincare            `json:"skincare"`
	AcneTpye   []entities.AcneFacial `json:"acne_type"`
	FacialType []entities.AcneFacial `json:"facial_type"`
	CreateAt   *time.Time            `json:"create_at"`
}

type Thread struct {
	ID           uint           `json:"id"`
	UserID       uint           `json:"user_id"`
	User         User           `json:"user"`
	Bookmark     bool           `json:"bookmark"`
	Favorite     bool           `json:"favorite"`
	ThreadDetail []ThreadDetail `json:"thread_detail"`
}

type ThreadDetail struct {
	ID       uint     `json:"id"`
	Skincare Skincare `json:"skincare"`
	Caption  string   `json:"caption"`
}

func DeleteResponse(id int) *Responses {
	return &Responses{
		Status: true,
		Data: map[string]string{
			"delete_id": string(rune(id)),
		},
		Error: nil,
	}
}

func ErrorResponse(err error) *Responses {
	return &Responses{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}

func TokenResponse(token string) *Responses {
	return &Responses{
		Status: true,
		Data: map[string]string{
			"token": "Bearer " + token,
		},
		Error: nil,
	}
}
