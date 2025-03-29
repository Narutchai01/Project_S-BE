package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicFacialResponse(data entities.FaceProblem) Facial {
	facial := Facial{
		ID:       data.ID,
		Name:     data.Name,
		Image:    data.Image,
		CreateBY: uint(data.CreatedBy),
	}
	return facial
}
func ToFacialResponse(data entities.FaceProblem) *Responses {
	facial := PublicFacialResponse(data)
	return &Responses{
		Status: true,
		Data:   facial,
		Error:  nil,
	}
}

func ToFacialsResponse(data []entities.FaceProblem) *Responses {
	facials := []Facial{}

	for _, facial := range data {
		facials = append(facials, PublicFacialResponse(facial))
	}
	return &Responses{
		Status: true,
		Data:   facials,
		Error:  nil,
	}
}
