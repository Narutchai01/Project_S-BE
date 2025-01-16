package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func ToAcneResponse(data entities.Acne) *Responses {
	acne := Acne{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: data.CreateBY,
	}
	return &Responses{
		Status: true,
		Data:   acne,
		Error:  nil,
	}
}

func ToAcnesResponse(data []entities.Acne) *Responses {
	acnes := []Acne{}

	for _, acne := range data {
		acnes = append(acnes, Acne{
			ID:       acne.ID,
			Name:     acne.Name,
			Image:    acne.Image,
			CreateBY: acne.CreateBY,
		})
	}
	return &Responses{
		Status: true,
		Data:   acnes,
		Error:  nil,
	}
}
