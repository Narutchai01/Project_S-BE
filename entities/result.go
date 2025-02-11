package entities

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	Image      string       `json:"image"`
	UserID     uint         `json:"user_id"`
	AcneType   []AcneFacial `json:"acne_type" gorm:"serializer:json"`
	FacialType []AcneFacial `json:"facial_type" gorm:"serializer:json"`
	SkinID     uint         `json:"skin_id"`
	SkincareID []uint       `json:"skincare_id" gorm:"serializer:json"`
	Skincare   []Skincare   `json:"skincare" gorm:"-"`
}

type AcneFacial struct {
	ID    uint `json:"id"`
	Count uint `json:"count"`
}
