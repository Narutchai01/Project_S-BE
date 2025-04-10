package entities

import "gorm.io/gorm"

type FavoriteCommentThread struct {
	gorm.Model
	CommentID uint `json:"comment_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	UserID    uint `json:"user_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	Status    bool `json:"status" gorm:"default:true"`
}

type FavoriteThread struct {
	gorm.Model
	ThreadID uint   `json:"thread_id" gorm:"not null"`
	Thread   Thread `gorm:"foreignKey:ThreadID"`
	UserID   uint   `json:"user_id" gorm:"not null;"`
	User     User   `gorm:"foreignKey:UserID"`
	Status   bool   `json:"status" gorm:"default:true"`
}

type FavoriteReviewSkincare struct {
	gorm.Model
	ReviewSkincareID uint           `json:"review_skincare_id" gorm:"not null;uniqueIndex:idx_review_skincare_user"`
	ReviewSkincare   ReviewSkincare `gorm:"foreignKey:ReviewSkincareID"`
	UserID           uint           `json:"user_id" gorm:"not null;uniqueIndex:idx_review_skincare_user"`
	User             User           `gorm:"foreignKey:UserID"`
	Status           bool           `json:"status" gorm:"default:true"`
}

type FavoriteCommentReviewSkincare struct {
	gorm.Model
	CommentID        uint `json:"comment_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	ReviewSkincareID uint `json:"review_skincare_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	UserID           uint `json:"user_id" gorm:"not null;uniqueIndex:idx_comment_user"`
	Status           bool `json:"status" gorm:"default:true"`
}

type Favorite struct {
	gorm.Model
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CommunityID uint      `json:"community_id" gorm:"default:null"`
	Community   Community `gorm:"foreignKey:CommunityID"`
	CommentID   uint      `json:"comment_id" gorm:"default:null"`
	Comment     Comment   `gorm:"foreignKey:CommentID"`
}
