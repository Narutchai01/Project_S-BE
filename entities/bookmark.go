package entities

import "gorm.io/gorm"

type BookmarkThread struct {
	gorm.Model `swaggerignore:"true"`
	ThreadID   uint   `json:"thread_id" gorm:"not null;index"`
	UserID     uint   `json:"user_id" gorm:"not null;index"`
	Thread     Thread `gorm:"foreignKey:ThreadID"`
	Status     bool   `json:"status" gorm:"default:true"`
}
type BookmarkReviewSkincare struct {
	gorm.Model       `swaggerignore:"true"`
	ReviewSkincareID uint           `json:"review_skincare_id" gorm:"not null;index"`
	UserID           uint           `json:"user_id" gorm:"not null;index"`
	ReviewSkincare   ReviewSkincare `gorm:"foreignKey:ReviewSkincareID"`
	Status           bool           `json:"status" gorm:"default:true"`
}
