package entities

import (
	"gorm.io/gorm"
)

type Skincare struct {
	gorm.Model
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreateBY uint `reqHeader:"create_by"`
}