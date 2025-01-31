package presentation

import (
	"strconv"
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
	SensitiveSkin bool       `json:"sensitive_skin"`
	Image         string     `json:"image"`
}

type Facial struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	CreateBY uint   `json:"create_by"`
}

type Result struct {
	ID         uint                          `json:"id"`
	Image      string                        `json:"image"`
	UserId     uint                          `json:"user_id"`
	AcneType   []entities.Acne_Facial_Result `json:"acne_type"`
	FacialType []entities.Acne_Facial_Result `json:"facial_type"`
	SkinType   uint                          `json:"skin_type"`
	Skincare   []uint                        `json:"skincare"`
}

func DeleteResponse(id int) *Responses {
	return &Responses{
		Status: true,
		Data: map[string]string{
			"delete_id": strconv.Itoa(id),
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
			"token": token,
		},
		Error: nil,
	}
}
