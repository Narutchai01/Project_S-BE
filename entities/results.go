package entities

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	Image      string       `json:"image"`
	UserID     uint         `json:"user_id"`
	AcneType   []AcneFacial `json:"acne_type" gorm:"json"`
	FacialType []AcneFacial `json:"facial_type" gorm:"json"`
	SkinID     uint         `json:"skin_id"`
	SkincareID []uint       `json:"skincare_id" gorm:"json"`
	Skincare   []Skincare   `json:"skincare" gorm:"many2many:skincare_result"`
}

type AcneFacial struct {
	ID    uint `json:"id"`
	Count uint `json:"count"`
}
