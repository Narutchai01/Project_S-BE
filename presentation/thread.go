package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func SkincareMap(data entities.Skincare) Skincare {
	skincare := Skincare{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Image:       data.Image,
		CreateBY:    data.CreateBY,
	}

	return skincare
}

func MapThreadDetails(data []entities.ThreadDetail) []ThreadDetail {
	threadDetail := []ThreadDetail{}

	for _, detail := range data {
		threadDetail = append(threadDetail, ThreadDetail{
			ID:       detail.ID,
			Skincare: SkincareMap(detail.Skincare),
			Caption:  detail.Caption,
		})
	}

	return threadDetail
}

func ToThreadResponse(data entities.Thread) *Responses {
	threads := Thread{
		ID:     data.ID,
		UserID: data.UserID,
		User: User{
			ID:       data.UserID,
			FullName: data.User.FullName,
			Email:    data.User.Email,
		},
		ThreadDetail: MapThreadDetails(data.Threads),
	}

	return &Responses{
		Status: true,
		Data:   threads,
		Error:  nil,
	}
}

func ToThreadListResponse(data []entities.Thread) *Responses {
	threads := []Thread{}

	for _, thread := range data {
		threads = append(threads, Thread{
			ID:     thread.ID,
			UserID: thread.UserID,
			User: User{
				ID:       thread.UserID,
				FullName: thread.User.FullName,
				Email:    thread.User.Email,
			},
			ThreadDetail: MapThreadDetails(thread.Threads),
		})
	}

	return &Responses{
		Status: true,
		Data:   threads,
		Error:  nil,
	}
}
