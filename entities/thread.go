package entities

import "gorm.io/gorm"

type Thread struct {
	gorm.Model `swaggerignore:"true"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	Title      string `json:"title" gorm:"not null"`
	Image      string `json:"image"`
	User       User   `gorm:"foreignKey:UserID;references:ID"`
	Bookmark   bool   `json:"bookmark" gorm:"-"`
	Owner      bool   `json:"owner" gorm:"-"`
	Favorite   bool   `json:"favorite"`
	Threads    []ThreadDetail
}

type ThreadDetail struct {
	gorm.Model `swaggerignore:"true"`
	ThreadID   uint     `json:"thread_id" swaggerignore:"true"`
	SkincareID uint     `json:"skincare_id"`
	Skincare   Skincare `swaggerignore:"true"`
	Caption    string   `json:"caption"`
}

type ThreadRequest struct {
	Title        string         `json:"title"`
	ThreadDetail []ThreadDetail `json:"thread_detail"`
}
