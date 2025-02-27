package entities

import "gorm.io/gorm"

type ReviewSkincare struct {
	gorm.Model    `swaggerignore:"true"`
	Title         string     `json:"title" gorm:"not null"`
	Content       string     `json:"content" gorm:"not null"`
	Favorite      bool       `json:"favorite" gorm:"-"`
	FavoriteCount int64      `json:"favorite_count" gorm:"-"`
	Bookmark      bool       `json:"bookmark" gorm:"-"`
	Owner         bool       `json:"owner" gorm:"-"`
	Image         string     `json:"image" swaggerignore:"true" gorm:"not null"`
	UserID        uint       `json:"user_id" gorm:"not null"`
	User          User       `gorm:"foreignKey:UserID;references:ID"`
	SkincareID    []int      `json:"skincare_id" gorm:"serializer:json"`
	Skincare      []Skincare `json:"skincare" gorm:"-"`
}
