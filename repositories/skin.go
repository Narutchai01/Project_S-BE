package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type SkinRepository interface {
	CreateSkin(skin entities.Skin) (entities.Skin, error)
	GetSkins() ([]entities.Skin, error)
	GetSkin(id int) (entities.Skin, error)
	UpdateSkin(id int, skin entities.Skin) (entities.Skin, error)
	DeleteSkin(id int) error
}
