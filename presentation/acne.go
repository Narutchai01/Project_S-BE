package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicAcneResponse(data entities.FaceProblem) Acne {
	acne := Acne{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: uint(data.CreatedBy),
	}
	return acne
}

func ToAcneResponse(data entities.FaceProblem) *Responses {
	acne := PublicAcneResponse(data)
	return &Responses{
		Status: true,
		Data:   acne,
		Error:  nil,
	}
}

func ToAcnesResponse(data []entities.FaceProblem) *Responses {
	acnes := []Acne{}

	for _, acne := range data {
		acnes = append(acnes, PublicAcneResponse(acne))
	}
	return &Responses{
		Status: true,
		Data:   acnes,
		Error:  nil,
	}
}
