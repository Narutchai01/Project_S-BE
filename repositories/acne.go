package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type AcneRepository interface {
	CreateAcne(acne entities.Acne) (entities.Acne, error)
}
