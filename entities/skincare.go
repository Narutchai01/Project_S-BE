package entities

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Skincare struct {
	gorm.Model  `swaggerignore:"true"`
	Image       string          `json:"image" swaggerignore:"true" gorm:"default:null"`
	Name        string          `json:"name" gorm:"not null unique"`
	Description string          `json:"description" gorm:"default:null"`
	CreateBY    uint            `reqHeader:"create_by" swaggerignore:"true"`
	Admin       Admin           `gorm:"foreignKey:CreateBY;references:ID"`
	Embedding   pgvector.Vector `gorm:"type:vector(384);default:null"`
}
