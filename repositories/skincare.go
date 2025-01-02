package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type SkincareRepository interface {
	CreateSkincare(skincare entities.Skincare) (entities.Skincare, error)
	GetSkincares() ([]entities.Skincare, error)
	GetSkincare(id int) (entities.Skincare, error)
	UpdateSkincare(id int, skincare entities.Skincare) (entities.Skincare, error)
	DeleteSkincare(id int) (entities.Skincare, error)
}