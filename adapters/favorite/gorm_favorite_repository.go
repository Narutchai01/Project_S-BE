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

func (repo *GormFavoriteRepository) FavoriteComment(comment_id uint, user_id uint) (entities.FavoriteComment, error) {
	favorite := entities.FavoriteComment{
		UserID:    user_id,
		CommentID: comment_id,
	}

	if err := repo.db.Create(&favorite).Error; err != nil {
		return entities.FavoriteComment{}, err
	}

	return favorite, nil
}

func (repo *GormFavoriteRepository) FindFavoriteComment(thread_id uint, user_id uint) (entities.FavoriteComment, error) {
	var favorite entities.FavoriteComment

	if err := repo.db.Where("comment_id = ? AND user_id = ?", thread_id, user_id).First(&favorite).Error; err != nil {
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
