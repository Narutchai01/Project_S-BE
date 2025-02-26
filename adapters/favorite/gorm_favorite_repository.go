package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormFavoriteRepository struct {
	db *gorm.DB
}

func NewGormFavoriteRepository(db *gorm.DB) repositories.FavoriteRepository {
	return &GormFavoriteRepository{db: db}
}

func (r *GormFavoriteRepository) FavoriteComment(commentID uint, userID uint) (entities.FavoriteComment, error) {
	favoriteComment := entities.FavoriteComment{
		CommentID: commentID,
		UserID:    userID,
		Status:    true,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favoriteComment).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entities.FavoriteComment{}, err
	}

	return favoriteComment, nil
}

func (repo *GormFavoriteRepository) FindFavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error) {
	var favorite entities.FavoriteComment

	if err := repo.db.Where("comment_id = ? AND user_id = ?", comment_id, user_id).First(&favorite).Error; err != nil {
		return entities.FavoriteComment{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) UpdateFavoriteComment(favorite_comment entities.FavoriteComment) (entities.FavoriteComment, error) {
	if err := repo.db.Save(&favorite_comment).Error; err != nil {
		return entities.FavoriteComment{}, err
	}

	return favorite_comment, nil
}

func (r *GormFavoriteRepository) FavoriteThread(threadID uint, userID uint) (entities.FavoriteThread, error) {
	favoriteThread := entities.FavoriteThread{
		ThreadID: threadID,
		UserID:   userID,
		Status:   true,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favoriteThread).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entities.FavoriteThread{}, err
	}

	return favoriteThread, nil
}

func (repo *GormFavoriteRepository) FindFavoriteThread(thread_id uint, user_id uint) (entities.FavoriteThread, error) {
	var favorite entities.FavoriteThread

	if err := repo.db.Where("thread_id = ? AND user_id = ?", thread_id, user_id).First(&favorite).Error; err != nil {
		return entities.FavoriteThread{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) UpdateFavoriteThread(favorite_thread entities.FavoriteThread) (entities.FavoriteThread, error) {
	if err := repo.db.Save(&favorite_thread).Error; err != nil {
		return entities.FavoriteThread{}, err
	}

	return favorite_thread, nil
}

func (repo *GormFavoriteRepository) CountFavoriteThread(thread_id uint) (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.FavoriteThread{}).Where("thread_id = ? AND status != false", thread_id).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
