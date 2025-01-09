package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
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
