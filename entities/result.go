package entities

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model `swaggerignore:"true"`
	Image      string               `json:"image" gorm:"not null"`
	UserId     uint                 `json:"user_id" gorm:"not null"`
	AcneType   []Acne_Facial_Result `json:"acne_type" gorm:"type:jsonb;serializer:json"`
	FacialType []Acne_Facial_Result `json:"facial_type" gorm:"type:jsonb;serializer:json"`
	SkinType   uint                 `json:"skin_type" gorm:"not null"`
	Skincare   []uint               `gorm:"type:jsonb;serializer:json" json:"skincare"`
}

type Acne_Facial_Result struct {
	ID    uint `json:"id"`
	Count uint `json:"count"`
}
