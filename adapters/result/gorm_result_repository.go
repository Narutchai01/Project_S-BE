package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormResultRepository struct {
	db *gorm.DB
}

func NewGormResultRepository(db *gorm.DB) repositories.ResultsRepository {
	return &GormResultRepository{db: db}
}

func (repo *GormResultRepository) CreateResult(result entities.Result) (entities.Result, error) {
	err := repo.db.Create(&result).Error
	if err != nil {
		return result, err
	}
	skincares, err2 := repo.FindSkincare(result.SkincareID)
	if err2 != nil {
		return result, err2
	}
	result.Skincare = skincares
	return result, nil
}

func (repo *GormResultRepository) FindSkincare(ids []uint) ([]entities.Skincare, error) {
	var skincares []entities.Skincare
	err := repo.db.Select("ID, image, name, description").Find(&skincares, ids).Error
	if err != nil {
		return nil, err
	}
	return skincares, nil
}

func (repo *GormResultRepository) GetResults(id uint) ([]entities.Result, error) {
	var results []entities.Result
	err := repo.db.Where("user_id = ?", id).Order("id asc").Find(&results).Error
	if err != nil {
		return nil, err
	}

	for i, result := range results {
		skincares, err := repo.FindSkincare(result.SkincareID)
		if err != nil {
			return nil, err
		}
		results[i].Skincare = skincares

		skin, err := repo.FindSkin(result.SkinID)
		if err != nil {
			return nil, err
		}
		results[i].Skin = skin
	}

	return results, nil
}

func (repo *GormResultRepository) GetResultLatest(id uint) (entities.Result, error) {
	var result entities.Result
	err := repo.db.Where("user_id = ?", id).Order("created_at desc").First(&result).Error
	if err != nil {
		return result, err
	}
	skincares, err2 := repo.FindSkincare(result.SkincareID)
	if err2 != nil {
		return result, err2
	}
	result.Skincare = skincares
	return result, nil
}

func (repo *GormResultRepository) GetResult(id uint) (entities.Result, error) {
	var result entities.Result
	err := repo.db.First(&result, id).Error
	if err != nil {
		return result, err
	}
	skincares, err2 := repo.FindSkincare(result.SkincareID)
	if err2 != nil {
		return result, err2
	}
	result.Skincare = skincares

	skin, err3 := repo.FindSkin(result.SkinID)
	if err3 != nil {
		return result, err3
	}
	result.Skin = skin

	return result, nil
}

func (repo *GormResultRepository) FindSkin(id uint) (entities.Skin, error) {
	var result entities.Skin
	err := repo.db.First(&result, id).Error
	return result, err
}
