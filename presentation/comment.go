package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicCommentThread(data entities.CommentThread) CommentThread {
	var commment = CommentThread{
		ID:            data.ID,
		ThreadID:      data.ThreadID,
		User:          *PublicUser(data.User),
		Favorite:      data.Favorite,
		FavoriteCount: data.FavoriteCount,
		Text:          data.Text,
	}
	return commment
}

func PublicCommentsThread(datas []entities.CommentThread) []CommentThread {
	var comments []CommentThread

	for _, comment := range datas {
		comments = append(comments, PublicCommentThread(comment))
	}

	return comments
}

func ToCommentThread(data entities.CommentThread) *Responses {

	comment := PublicCommentThread(data)
	return &Responses{
		Status: true,
		Data:   comment,
		Error:  nil,
	}
}

func ToCommentsThread(datas []entities.CommentThread) *Responses {
	comments := PublicCommentsThread(datas)

	return &Responses{
		Status: true,
		Data:   comments,
		Error:  nil,
	}
}

func PublicCommentReviewSkincare(data entities.CommentReviewSkicare) CommentReviewSkicare {
	var commment = CommentReviewSkicare{
		ID:               data.ID,
		ReviewSkincareID: data.ReviewSkincareID,
		User:             *PublicUser(data.User),
		Favorite:         data.Favorite,
		FavoriteCount:    data.FavoriteCount,
		Content:          data.Content,
		CreateAt:         data.CreatedAt,
	}
	return commment
}

func PublicCommentsReviewSkincare(datas []entities.CommentReviewSkicare) []CommentReviewSkicare {
	var comments []CommentReviewSkicare

	for _, comment := range datas {
		comments = append(comments, PublicCommentReviewSkincare(comment))
	}

	return comments
}

func ToCommentReviewSkincare(data entities.CommentReviewSkicare) *Responses {

	comment := PublicCommentReviewSkincare(data)
	return &Responses{
		Status: true,
		Data:   comment,
		Error:  nil,
	}
}

func ToCommentsReviewSkincare(datas []entities.CommentReviewSkicare) *Responses {

	comments := PublicCommentsReviewSkincare(datas)

	return &Responses{
		Status: true,
		Data:   comments,
		Error:  nil,
	}
}
