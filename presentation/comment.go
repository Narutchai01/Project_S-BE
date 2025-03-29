package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicCommentThread(data entities.Comment) CommentThread {
	var commment = CommentThread{
		ID:            data.ID,
		ThreadID:      data.CommunityID,
		User:          *PublicUser(data.User),
		Favorite:      data.Favorite,
		FavoriteCount: data.FavoriteCount,
		Text:          data.Content,
		CreateAt:      data.CreatedAt,
	}
	return commment
}

func PublicCommentsThread(datas []entities.Comment) []CommentThread {
	var comments []CommentThread

	for _, comment := range datas {
		comments = append(comments, PublicCommentThread(comment))
	}

	return comments
}

func ToCommentThread(data entities.Comment) *Responses {

	comment := PublicCommentThread(data)
	return &Responses{
		Status: true,
		Data:   comment,
		Error:  nil,
	}
}

func ToCommentsThread(datas []entities.Comment) *Responses {
	comments := PublicCommentsThread(datas)

	return &Responses{
		Status: true,
		Data:   comments,
		Error:  nil,
	}
}

func PublicCommentReviewSkincare(data entities.Comment) CommentReviewSkicare {
	var commment = CommentReviewSkicare{
		ID:               data.ID,
		ReviewSkincareID: data.CommunityID,
		User:             *PublicUser(data.User),
		// Favorite:         data.Favorite,
		// FavoriteCount:    data.FavoriteCount,
		Content:  data.Content,
		CreateAt: data.CreatedAt,
	}
	return commment
}

func PublicCommentsReviewSkincare(datas []entities.Comment) []CommentReviewSkicare {
	var comments []CommentReviewSkicare

	for _, comment := range datas {
		comments = append(comments, PublicCommentReviewSkincare(comment))
	}

	return comments
}

func ToCommentReviewSkincare(data entities.Comment) *Responses {

	comment := PublicCommentReviewSkincare(data)
	return &Responses{
		Status: true,
		Data:   comment,
		Error:  nil,
	}
}

func ToCommentsReviewSkincare(datas []entities.Comment) *Responses {

	comments := PublicCommentsReviewSkincare(datas)

	return &Responses{
		Status: true,
		Data:   comments,
		Error:  nil,
	}
}
