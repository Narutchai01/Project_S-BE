package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"gorm.io/gorm"
)

type GormResultRepository struct {
	db *gorm.DB
}

func NewGormResultRepository(db *gorm.DB) repositories.ResultRepository {
	return &GormResultRepository{db: db}
}

var query = `SELECT
	results.id,
	results.image,
	results.user_id,
	results.acne_type,
	results.facial_type,
	results.skin_type,
	jsonb_agg(
      	jsonb_build_object(
			'id', skincares.id,
			'image', skincares.image,
			'name', skincares.name,
			'description', skincares.description,
			'create_by', skincares.create_by
      	)
	) AS skincare
	FROM results
		JOIN LATERAL jsonb_array_elements_text(results.skincare :: jsonb) AS skincare_id ON TRUE
		JOIN skincares ON skincares.id = skincare_id :: INT
	WHERE (results.id = ? OR ? = 0) AND (results.user_id = ? OR ? = 0) AND results.deleted_at IS NULL
	GROUP BY
		results.id,
		results.image,
		results.user_id,
		results.acne_type,
		results.facial_type,
		results.skin_type
	ORDER BY results.id DESC
`

func (repo *GormResultRepository) CreateResult(result entities.Result) (entities.ResultWithSkincare, error) {
	if err := repo.db.Create(&result).Error; err != nil {
		return entities.ResultWithSkincare{}, err
	}
	var resultWithSkincare entities.ResultWithSkincare

	err := repo.db.Raw(query, 0, 0, 0, 0).First(&resultWithSkincare).Error
	return resultWithSkincare, err
}

func (repo *GormResultRepository) GetResults() ([]entities.ResultWithSkincare, error) {
	var results []entities.ResultWithSkincare
	err := repo.db.Raw(query, 0, 0, 0, 0).Scan(&results).Error
	return results, err
}

func (repo *GormResultRepository) GetResultById(id int) (entities.ResultWithSkincare, error) {
	var result entities.ResultWithSkincare
	err := repo.db.Raw(query, id, id, 0, 0).First(&result).Error
	return result, err
}

func (repo *GormResultRepository) GetResultByIdWithoutSkincare(id int) (entities.Result, error) {
	var result entities.Result
	err := repo.db.First(&result, id, 0, 0).Error
	return result, err
}

func (repo *GormResultRepository) UpdateResultById(id int, result entities.Result) (entities.ResultWithSkincare, error) {
	if err := repo.db.Model(&entities.Result{}).Where("id = ?", id).Updates(&result).Error; err != nil {
		return entities.ResultWithSkincare{}, err
	}

	var updateResult entities.ResultWithSkincare
	err := repo.db.Raw(query, id, id, 0, 0).First(&updateResult).Error
	return updateResult, err
}

func (repo *GormResultRepository) DeleteResultById(id int) error {
	err := repo.db.Where("id = ?", id).Delete(&entities.Result{}).Error
	return err
}

func (repo *GormResultRepository) GetResultsByUserId(user_id int) ([]entities.ResultWithSkincare, error) {
	var results []entities.ResultWithSkincare
	err := repo.db.Raw(query, 0, 0, user_id, user_id).Scan(&results).Error
	return results, err
}

func (repo *GormResultRepository) GetLatestResultByUserIdFromToken(user_id int) (entities.ResultWithSkincare, error) {
	var result entities.ResultWithSkincare
	err := repo.db.Raw(query, 0, 0, user_id, user_id).Last(&result).Error
	return result, err
}
