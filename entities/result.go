package entities

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	Image      string           `json:"image"`
	UserID     uint             `json:"user_id"`
	AcneType   []AcneFacial     `json:"acne_type" gorm:"serializer:json"`
	FacialType []AcneFacial     `json:"facial_type" gorm:"serializer:json"`
	SkinID     uint             `json:"skin_id"`
	Skin       FaceProblem      `json:"skin" gorm:"foreignKey:SkinID"`
	SkincareID []uint           `json:"skincare_id" gorm:"-"`
	Skincare   []SkincareResult `json:"skincare" gorm:"foreignKey:ResultID"`
	User       User             `json:"user" gorm:"foreignKey:UserID"`
}

type AcneFacial struct {
	ID    uint `json:"id"`
	Count uint `json:"count"`
}

type SkincareResult struct {
	gorm.Model
	SkincareID uint     `json:"skincare_id"`
	Skincare   Skincare `json:"skincare" gorm:"foreignKey:SkincareID"`
	ResultID   uint     `json:"result_id"`
}
