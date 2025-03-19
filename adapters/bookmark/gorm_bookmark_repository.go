package adapters

import (
	"errors"

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

func (repo *GormBookmarkRepository) CreateBookmarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error) {
	bookmark := entities.BookmarkThread{
		ThreadID: thread_id,
		UserID:   user_id,
	}
	if err := repo.db.Create(&bookmark).Error; err != nil {
		return entities.BookmarkThread{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) FindBookMarkThread(thread_id uint, user_id uint) (entities.BookmarkThread, error) {
	var bookmark entities.BookmarkThread
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.BookmarkThread{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) UpdateBookMarkThread(thread_id uint, user_id uint, status bool) (entities.BookmarkThread, error) {
	var bookmark entities.BookmarkThread
	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&bookmark).Error; err != nil {
		return entities.BookmarkThread{}, err
	}

	bookmark.Status = status
	if err := repo.db.Save(&bookmark).Error; err != nil {
		return entities.BookmarkThread{}, err
	}

	return bookmark, nil
}

func (repo *GormBookmarkRepository) BookmarkReviewSkincare(review_id uint, user_id uint) (entities.BookmarkReviewSkincare, error) {
	bookmark := entities.BookmarkReviewSkincare{
		ReviewSkincareID: review_id,
		UserID:           user_id,
	}
	if err := repo.db.Create(&bookmark).Error; err != nil {
		return entities.BookmarkReviewSkincare{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) FindBookMarkReviewSkincare(review_id uint, user_id uint) (entities.BookmarkReviewSkincare, error) {
	var bookmark entities.BookmarkReviewSkincare
	if err := repo.db.Where("review_skincare_id = ? AND user_id = ?", review_id, user_id).First(&bookmark).Error; err != nil {
		return entities.BookmarkReviewSkincare{}, err
	}
	return bookmark, nil
}

func (repo *GormBookmarkRepository) UpdateBookMarkReviewSkincare(review_id uint, user_id uint, status bool) (entities.BookmarkReviewSkincare, error) {
	var bookmark entities.BookmarkReviewSkincare
	if err := repo.db.Where("review_skincare_id = ? AND user_id = ?", review_id, user_id).First(&bookmark).Error; err != nil {
		return entities.BookmarkReviewSkincare{}, err
	}

	bookmark.Status = status
	if err := repo.db.Save(&bookmark).Error; err != nil {
		return entities.BookmarkReviewSkincare{}, err
	}

	return bookmark, nil
}

func (repo *GormBookmarkRepository) FindBookmark(community_id uint, user_id uint) (bool, entities.Bookmark, error) {
	var bookmark entities.Bookmark

	if err := repo.db.Where("community_id = ? and user_id = ?", community_id, user_id).First(&bookmark).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, entities.Bookmark{}, err
		}
		return false, entities.Bookmark{}, err
	}
	return true, bookmark, nil
}

func (repo *GormBookmarkRepository) Bookmark(community_id uint, user_id uint) (entities.Bookmark, error) {

	bookmark := entities.Bookmark{
		CommunityID: community_id,
		UserID:      user_id,
	}

	if err := repo.db.Create(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	return bookmark, nil

}

func (repo *GormBookmarkRepository) DeleteBookmark(community_id uint, user_id uint) (entities.Bookmark, error) {
	var bookmark entities.Bookmark

	if err := repo.db.Where("community_id = ? and user_id = ?", community_id, user_id).Delete(&bookmark).Error; err != nil {
		return entities.Bookmark{}, err
	}

	return bookmark, nil
}
