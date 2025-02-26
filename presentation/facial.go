package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func ToFacialResponse(data entities.Facial) *Responses {
	facial := Facial{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: data.CreateBY,
	}
	return &Responses{
		Status: true,
		Data:   facial,
		Error:  nil,
	}
}

func ToFacialsResponse(data []entities.Facial) *Responses {
	facials := []Facial{}

	for _, facial := range data {
		facials = append(facials, Facial{
			ID:       facial.ID,
			Name:     facial.Name,
			Image:    facial.Image,
			CreateBY: facial.CreateBY,
		})
	}
	return &Responses{
		Status: true,
		Data:   facials,
		Error:  nil,
	}
}
