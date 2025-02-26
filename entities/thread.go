package entities

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Title         string        `json:"title"`
	Caption       string        `json:"caption"`
	UserID        uint          `json:"user_id"`
	User          User          `json:"user"`
	Favorite      bool          `json:"favorite" gorm:"-"`
	FavoriteCount int64         `json:"favorite_count" gorm:"-"`
	Bookmark      bool          `json:"bookmark" gorm:"-"`
	Images        []ThreadImage `json:"images" gorm:"-"`
}

type ThreadImage struct {
	gorm.Model
	ThreadID uint   `json:"thread_id"`
	Thread   Thread `gorm:"foreignKey:ThreadID"`
	Image    string `json:"image"`
}
