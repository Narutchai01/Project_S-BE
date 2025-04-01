package entities

import "gorm.io/gorm"

type CommunityType struct {
	gorm.Model
	Type string `json:"type"`
}

type Community struct {
	gorm.Model
	Title      string              `json:"title"`
	Caption    string              `json:"caption"`
	Owner      bool                `json:"owner" gorm:"-"`
	Likes      uint64              `json:"likes" gorm:"default:0"`
	Favorite   bool                `json:"favorite" gorm:"-"`
	Bookmark   bool                `json:"bookmark" gorm:"-"`
	TypeID     uint64              `json:"type_id" gorm:"not null"`
	Type       CommunityType       `json:"type" gorm:"foreignKey:TypeID"`
	UserID     uint64              `json:"user_id" gorm:"not null"`
	User       User                `json:"user" gorm:"foreignKey:UserID"`
	Images     []CommunityImage    `json:"images" gorm:"foreignKey:CommunityID;references:ID"`
	SkincareID []int               `json:"skincare_id" gorm:"-"`
	Skincares  []SkincareCommunity `json:"skincares" gorm:"foreignKey:CommunityID;references:ID"`
}

type CommunityImage struct {
	gorm.Model
	Image       string `json:"image"`
	CommunityID uint64 `json:"community_id"`
}

type SkincareCommunity struct {
	gorm.Model
	SkincareID  uint64   `json:"skincare_id"`
	Skincare    Skincare `json:"skincare" gorm:"foreignKey:SkincareID"`
	CommunityID uint64   `json:"community_id"`
}
