package entities

import (
	"gorm.io/gorm"
)

type Skincare struct {
	gorm.Model  `swaggerignore:"true"`
	Image       string `json:"image" swaggerignore:"true"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateBY    uint   `reqHeader:"create_by" swaggerignore:"true"`
}
