package entities

import "gorm.io/gorm"

type CommentThread struct {
	gorm.Model
	ThreadID      uint   `json:"thread_id"`
	UserID        uint   `json:"user_id"`
	User          User   `json:"user" gorm:"foreignKey:UserID"`
	Favorite      bool   `json:"favorite"`
	FavoriteCount int    `json:"favorite_count" gorm:"-"`
	Text          string `json:"text"`
}

type CommentReviewSkicare struct {
	gorm.Model
	ReviewSkincareID uint           `json:"review_skincare_id"`
	ReviewSkincare   ReviewSkincare `json:"review_skincare" gorm:"foreignKey:ReviewSkincareID"`
	UserID           uint           `json:"user_id"`
	User             User           `json:"user" gorm:"foreignKey:UserID"`
	Favorite         bool           `json:"favorite" gorm:"-"`
	FavoriteCount    int            `json:"favorite_count" gorm:"-"`
	Content          string         `json:"content"`
}
