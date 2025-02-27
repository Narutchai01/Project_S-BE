package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicBookmarkThread(bookmark entities.BookmarkThread) BookmarkThread {
	return BookmarkThread{
		ID:       bookmark.ID,
		ThreadID: bookmark.ThreadID,
		UserID:   bookmark.UserID,
		Status:   bookmark.Status,
	}
}

func PublicBookmarkReviewSkincare(bookmark entities.BookmarkReviewSkincare) BookmarkReviewSkincare {
	return BookmarkReviewSkincare{
		ID:               bookmark.ID,
		ReviewSkincareID: bookmark.ReviewSkincareID,
		UserID:           bookmark.UserID,
		Status:           bookmark.Status,
	}
}

func ToBookmarkThreadResponse(data entities.BookmarkThread) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicBookmarkThread(data),
		Error:  nil,
	}
}

func ToBookmarkReviewSkincareResponse(data entities.BookmarkReviewSkincare) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicBookmarkReviewSkincare(data),
		Error:  nil,
	}
}
