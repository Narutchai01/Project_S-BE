package entities

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Title   string        `json:"title"`
	Caption string        `json:"caption"`
	UserID  uint          `json:"user_id"`
	User    User          `json:"user"`
	Images  []ThreadImage `json:"images" gorm:"-"`
}

type ThreadImage struct {
	gorm.Model
	ThreadID uint   `json:"thread_id"`
	Image    string `json:"image"`
}
