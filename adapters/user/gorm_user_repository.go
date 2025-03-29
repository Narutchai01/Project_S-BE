package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) CreateUser(user entities.User) (entities.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *GormUserRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repo *GormUserRepository) UpdateUserPasswordById(id int, newPassword string) (entities.User, error) {
	var user entities.User
	err := repo.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Update("password", newPassword).Error
	return user, err
}

func (repo *GormUserRepository) GetUser(id uint) (entities.User, error) {
	var user entities.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (repo *GormUserRepository) UpdateUser(user entities.User) (entities.User, error) {
	err := repo.db.Save(&user).Error
	return user, err
}

func (repo *GormUserRepository) Follower(follow_id uint, user_id uint) (entities.Follower, error) {
	var follower entities.Follower
	err := repo.db.Create(&entities.Follower{FollowerID: follow_id, UserID: user_id}).Error
	if err != nil {
		return follower, err
	}
	return follower, nil
}

func (repo *GormUserRepository) FindFollower(follow_id uint, user_id uint) (entities.Follower, error) {
	var follower entities.Follower
	err := repo.db.Where("follower_id = ? AND user_id = ?", follow_id, user_id).First(&follower).Error
	if err != nil {
		return follower, err
	}
	return follower, nil
}

func (repo *GormUserRepository) DeleteFollower(id uint) (entities.Follower, error) {
	var follower entities.Follower
	err := repo.db.Where("id = ?", id).Delete(&follower).Error
	if err != nil {
		return follower, err
	}
	return follower, nil
}
