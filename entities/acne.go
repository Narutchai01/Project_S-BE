package entities

import "gorm.io/gorm"

type Acne struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name"`
	Image      string `json:"image" swaggerignore:"true"`
	CreateBY   uint   `json:"create_by"`
}
