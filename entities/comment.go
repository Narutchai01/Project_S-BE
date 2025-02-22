package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ThreadID uint   `json:"thread_id"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Favorite bool   `json:"favorite"`
	Text     string `json:"text"`
}
