package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormBookmarkRepository struct {
	db *gorm.DB
}

func NewGormBookmarkRepository(db *gorm.DB) repositories.BookmarkRepository {
	return &GormBookmarkRepository{db: db}
}

func (repo *GormBookmarkRepository) CreateBookmark(thread_id uint, user_id uint) (entities.Bookmark, error) {
	bookmark := entities.Bookmark{
		ThreadID: thread_id,
		UserID:   user_id,
	}
	if err := repo.db.Create(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) FindBookMark(thread_id uint, user_id uint) (entities.Bookmark, error) {
	var bookmark entities.Bookmark
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) UpdateBookMark(thread_id uint, user_id uint, status bool) (entities.Bookmark, error) {
	var bookmark entities.Bookmark
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	bookmark.Status = status
	if err := repo.db.Save(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	return bookmark, nil
}
