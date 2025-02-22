package entities

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model `swaggerignore:"true"`
	ThreadID   uint   `json:"thread_id"`
	Thread     Thread `gorm:"foreignKey:ThreadID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true"`
	UserID     uint   `json:"user_id"`
	User       User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true"`
	Status     bool   `json:"status" gorm:"default:true"`
}
