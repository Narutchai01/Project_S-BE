package entities

import "gorm.io/gorm"

type FavoriteComment struct {
	gorm.Model
	CommentID uint `json:"comment_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	UserID    uint `json:"user_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	Status    bool `json:"status" gorm:"default:true"`
}
