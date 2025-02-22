package entities

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model `swaggerignore:"true"`
	ThreadID   uint `json:"thread_id" gorm:"not null;index"`
	UserID     uint `json:"user_id" gorm:"not null;index"`
	Status     bool `json:"status" gorm:"default:true"`
}
