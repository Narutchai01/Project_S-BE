package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type AcneRepository interface {
	CreateAcne(acne entities.Acne) (entities.Acne, error)
	GetAcnes() ([]entities.Acne, error)
	GetAcne(id int) (entities.Acne, error)
}
