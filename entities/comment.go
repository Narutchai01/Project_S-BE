package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ThreadID uint   `json:"thread_id"`
	UserID   uint   `json:"user_id"`
	Text     string `json:"text"`
}
