package entities

import "gorm.io/gorm"

type Acne struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" gorm:"not null unique"`
	Image      string `json:"image" swaggerignore:"true" gorm:"not null"`
	CreateBY   uint   `json:"create_by" swaggerignore:"true" gorm:"not null"`
}
