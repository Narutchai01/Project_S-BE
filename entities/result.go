package entities

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Image      string               `json:"image" gorm:"not null"`
	UserId     uint                 `json:"user_id" gorm:"not null"`
	AcneType   []Acne_Facial_Result `json:"acne_type" gorm:"type:bytes;serializer:gob"`
	FacialType []Acne_Facial_Result `json:"facial_type" gorm:"type:bytes;serializer:gob"`
	SkinType   uint                 `json:"skin_type" gorm:"not null"`
	Skincare   []uint               `gorm:"serializer:json" json:"skincare"`
}

type Acne_Facial_Result struct {
	ID    uint `json:"id" swaggerignore:"true"`
	Count uint `json:"count"`
}
