package entities

import "gorm.io/gorm"

type Skin struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" gorm:"not null unique"`
	Image      string `json:"image" swaggerignore:"true"`
	CreateBY   uint   `json:"create_by" swaggerignore:"true"`
	Admin      Admin  `gorm:"foreignKey:CreateBY;references:ID"`
}
