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
	ID            uint          `json:"id"`
	User          User          `json:"user"`
	Title         string        `json:"title"`
	Favorite      bool          `json:"favorite"`
	FavoriteCount int64         `json:"favorite_count"`
	Owner         bool          `json:"owner"`
	Bookmark      bool          `json:"bookmark"`
	Images        []ThreadImage `json:"images"`
	Caption       string        `json:"caption"`
	CreateAt      time.Time     `json:"create_at"`
}

type ThreadImage struct {
	ID       uint   `json:"id"`
	ThreadID uint   `json:"thread_id"`
	Image    string `json:"image"`
}

type ReviewSkincare struct {
	ID            uint       `json:"id"`
	Title         string     `json:"title"`
	Favortie      bool       `json:"favorite"`
	FavoriteCount int64      `json:"favorite_count"`
	Bookmark      bool       `json:"bookmark"`
	Owner         bool       `json:"owner"`
	Content       string     `json:"content"`
	Image         string     `json:"image"`
	User          User       `json:"user"`
	Skincare      []Skincare `json:"skincares"`
	CreateAt      time.Time  `json:"create_at"`
}
type BookmarkThread struct {
	ID       uint `json:"id"`
	ThreadID uint `json:"thread_id" `
	UserID   uint `json:"user_id" `
	Status   bool `json:"status" `
}
type BookmarkReviewSkincare struct {
	ID               uint `json:"id"`
	ReviewSkincareID uint `json:"review_skincare_id" `
	UserID           uint `json:"user_id" `
	Status           bool `json:"status" `
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
