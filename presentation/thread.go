package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicThreadImage(threadImages []entities.CommunityImage) []CommunityImage {
	var images []CommunityImage
	for _, image := range threadImages {
		images = append(images, CommunityImage{
			ID:          image.ID,
			CommunityID: uint(image.CommunityID),
			Image:       image.Image,
		})
	}

	return images
}

func PublicThread(thread entities.Community) Thread {
	return Thread{
		ID:            thread.ID,
		Title:         thread.Title,
		Owner:         thread.Owner,
		Favorite:      thread.Favorite,
		FavoriteCount: int64(thread.Likes),
		User:          *PublicUser(thread.User),

		Caption:  thread.Caption,
		Images:   PublicThreadImage(thread.Images),
		CreateAt: thread.CreatedAt,
	}
}

func ToThreadResponse(data entities.Community) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicThread(data),
		Error:  nil,
	}
}

func ToThreadsResponse(data []entities.Community) *Responses {
	var threads []Thread
	for _, thread := range data {
		threads = append(threads, PublicThread(thread))
	}

	return &Responses{
		Status: true,
		Data:   threads,
		Error:  nil,
	}
}
