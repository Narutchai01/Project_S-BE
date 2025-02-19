package entities

import "gorm.io/gorm"

type Thread struct {
	gorm.Model `swaggerignore:"true"`
	UserID     uint `json:"user_id" gorm:"not null"`
	User       User `gorm:"foreignKey:UserID;references:ID"`
	Threads    []ThreadDetail
}

type ThreadDetail struct {
	gorm.Model `swaggerignore:"true"`
	ThreadID   uint `json:"thread_id"`
	SkincareID uint `json:"skincare_id"`
	Skincare   Skincare
	Caption    string `json:"caption"`
}

type ThreadRequest struct {
	ThreadDetail []ThreadDetail `json:"thread_detail"`
}
