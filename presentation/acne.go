package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func ToAcneResponse(data entities.Acne) *Responses {
	acne := Acne{
		ID:    data.ID,
		Name:  data.Name,
		Image: data.Image,
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
			ID:    acne.ID,
			Name:  acne.Name,
			Image: acne.Image,
		})
	}
	return &Responses{
		Status: true,
		Data:   acnes,
		Error:  nil,
	}
}

func AcneErrorResponse(err error) *Responses {
	return &Responses{
		Status: false,
		Data:   nil,
		Error:  err.Error(),
	}
}

func DeleteAcneResponse(id int) *Responses {
	return &Responses{
		Status: true,
		Data:   map[string]string{"delete_id": string(rune(id))},
		Error:  nil,
	}
}
