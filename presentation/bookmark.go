package presentation

import "github.com/Narutchai01/Project_S-BE/entities"

func PublicBookmarkThread(bookmark entities.Bookmark) BookmarkThread {
	return BookmarkThread{
		ID:       bookmark.ID,
		ThreadID: bookmark.CommunityID,
		UserID:   bookmark.UserID,
		Status:   bookmark.Status,
	}
}

func PublicBookmarkReviewSkincare(bookmark entities.Bookmark) BookmarkReviewSkincare {
	return BookmarkReviewSkincare{
		ID:               bookmark.ID,
		ReviewSkincareID: bookmark.CommunityID,
		UserID:           bookmark.UserID,
		Status:           bookmark.Status,
	}
}

func ToBookmarkThreadResponse(data entities.Bookmark) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicBookmarkThread(data),
		Error:  nil,
	}
}

func ToBookmarkReviewSkincareResponse(data entities.Bookmark) *Responses {
	return &Responses{
		Status: true,
		Data:   PublicBookmarkReviewSkincare(data),
		Error:  nil,
	}
}
