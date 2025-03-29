package entities

import "gorm.io/gorm"

type FaceProblemType struct {
	gorm.Model
	Name string `json:"name"`
}

type FaceProblem struct {
	gorm.Model
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	TypeID    uint64          `json:"type_id" gorm:"not null"`
	Type      FaceProblemType `json:"type" gorm:"foreignKey:TypeID"`
	CreatedBy uint64          `json:"created_by" gorm:"not null"`
	Admin     Admin           `json:"admin" gorm:"foreignKey:CreatedBy"`
}
