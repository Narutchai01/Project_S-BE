package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type FacialRepository interface {
	CreateFacial(facial entities.Facial) (entities.Facial, error)
	GetFacials() ([]entities.Facial, error)
}
