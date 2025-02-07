package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type SkincareRepository interface {
	CreateSkincare(skincare entities.Skincare) (entities.Skincare, error)
	GetSkincares() ([]entities.Skincare, error)
	GetSkincareById(id int) (entities.Skincare, error)
	UpdateSkincareById(id int, skincare entities.Skincare) (entities.Skincare, error)
	DeleteSkincareById(id int) (entities.Skincare, error)
}
