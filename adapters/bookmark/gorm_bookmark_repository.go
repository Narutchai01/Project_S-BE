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
