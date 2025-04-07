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

func PublicCommunitiesBookmark(bookmarks entities.Bookmark) BookmarkCommunity {
	communityBokkmark := BookmarkCommunity{
		ID:          bookmarks.ID,
		CommunityID: bookmarks.CommunityID,
		UserID:      bookmarks.Community.User.ID,
		User:        *PublicUser(bookmarks.Community.User),
		Image:       bookmarks.Community.Images[0].Image,
		Title:       bookmarks.Community.Title,
		Content:     bookmarks.Community.Caption,
		Type:        int(bookmarks.Community.TypeID),
		Favorite:    bookmarks.Community.Favorite,
	}
	return communityBokkmark
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

func ToCommunityBookmarkResponse(data []entities.Bookmark) *Responses {
	bookmarks := []BookmarkCommunity{}

	if len(data) == 0 {
		return &Responses{
			Status: true,
			Data:   []BookmarkCommunity{},
			Error:  nil,
		}
	}

	for _, bookmark := range data {

		if len(bookmark.Community.Images) == 0 {
			continue
		}

		bookmarks = append(bookmarks, PublicCommunitiesBookmark(bookmark))
	}
	return &Responses{
		Status: true,
		Data:   bookmarks,
		Error:  nil,
	}
}
