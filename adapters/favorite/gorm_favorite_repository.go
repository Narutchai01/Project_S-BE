package adapters

import (
	"errors"
	"fmt"

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

func (r *GormFavoriteRepository) FavoriteCommentThread(commentID uint, userID uint) (entities.FavoriteCommentThread, error) {
	favoriteComment := entities.FavoriteCommentThread{
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
		return entities.FavoriteCommentThread{}, err
	}

	return favoriteComment, nil
}

func (repo *GormFavoriteRepository) FindFavoriteCommentThread(comment_id uint, user_id uint) (entities.FavoriteCommentThread, error) {
	var favorite entities.FavoriteCommentThread

	if err := repo.db.Where("comment_id = ? AND user_id = ?", comment_id, user_id).First(&favorite).Error; err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) UpdateFavoriteCommentThread(favorite_comment entities.FavoriteCommentThread) (entities.FavoriteCommentThread, error) {
	if err := repo.db.Save(&favorite_comment).Error; err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	return favorite_comment, nil
}

func (repo *GormFavoriteRepository) CountFavoriteCommentThread(comment_id uint) (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.FavoriteCommentThread{}).Where("comment_id = ? AND status != false", comment_id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
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

func (r *GormFavoriteRepository) FavoriteReviewSkincare(reviewSkincareID uint, userID uint) (entities.FavoriteReviewSkincare, error) {
	favoriteReviewSkincare := entities.FavoriteReviewSkincare{
		ReviewSkincareID: reviewSkincareID,
		UserID:           userID,
		Status:           true,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favoriteReviewSkincare).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	return favoriteReviewSkincare, nil
}

func (repo *GormFavoriteRepository) FindFavoriteReviewSkincare(review_skincare_id uint, user_id uint) (entities.FavoriteReviewSkincare, error) {
	var favorite entities.FavoriteReviewSkincare

	if err := repo.db.Where("review_skincare_id = ? AND user_id = ?", review_skincare_id, user_id).First(&favorite).Error; err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) UpdateFavoriteReviewSkincare(favorite_review_skincare entities.FavoriteReviewSkincare) (entities.FavoriteReviewSkincare, error) {
	if err := repo.db.Save(&favorite_review_skincare).Error; err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	return favorite_review_skincare, nil
}

func (repo *GormFavoriteRepository) CountFavoriteReviewSkincare(review_skincare_id uint) (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.FavoriteReviewSkincare{}).Where("review_skincare_id = ? AND status != false", review_skincare_id).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *GormFavoriteRepository) FavoriteCommentReviewSkincare(commentID uint, userID uint) (entities.FavoriteCommentReviewSkincare, error) {
	favoriteCommentReviewSkincare := entities.FavoriteCommentReviewSkincare{
		CommentID: commentID,
		UserID:    userID,
		Status:    true,
	}

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favoriteCommentReviewSkincare).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	return favoriteCommentReviewSkincare, nil
}

func (repo *GormFavoriteRepository) FindFavoriteCommentReviewSkincare(comment_id uint, user_id uint) (entities.FavoriteCommentReviewSkincare, error) {
	var favorite entities.FavoriteCommentReviewSkincare

	if err := repo.db.Where("comment_id = ? AND user_id = ?", comment_id, user_id).First(&favorite).Error; err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) UpdateFavoriteCommentReviewSkincare(favorite_comment_review_skincare entities.FavoriteCommentReviewSkincare) (entities.FavoriteCommentReviewSkincare, error) {
	if err := repo.db.Save(&favorite_comment_review_skincare).Error; err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	return favorite_comment_review_skincare, nil
}

func (repo *GormFavoriteRepository) CountFavoriteCommentReviewSkincare(comment_id uint) (int64, error) {
	var count int64
	if err := repo.db.Model(&entities.FavoriteCommentReviewSkincare{}).Where("comment_id = ? AND status != false", comment_id).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *GormFavoriteRepository) Favorite(favorite entities.Favorite) (entities.Favorite, error) {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favorite).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entities.Favorite{}, err
	}

	return favorite, nil
}

func (r *GormFavoriteRepository) FindFavorite(id uint, column string, user_id uint) (bool, uint, error) {
	var favorite entities.Favorite

	query := fmt.Sprintf("%s = ? and user_id = ?", column)
	if err := r.db.Where(query, id, user_id).First(&favorite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, 0, nil
		}
		return false, 0, err
	}

	return true, favorite.ID, nil
}

func (r *GormFavoriteRepository) DeleteFavorite(id uint) (entities.Favorite, error) {
	var favorite entities.Favorite

	if err := r.db.Where("id = ?", id).Delete(&favorite).Error; err != nil {
		return entities.Favorite{}, err
	}

	return favorite, nil
}

func (r *GormFavoriteRepository) CountFavorite(id uint, column string) int64 {
	var count int64
	query := fmt.Sprintf("%s = ? ", column)
	if err := r.db.Model(&entities.Favorite{}).Where(fmt.Sprintf(query, id), id).Count(&count).Error; err != nil {
		return 0
	}

	return count
}
