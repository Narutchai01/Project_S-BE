package usecases

import (
	"errors"
	"strings"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/repositories"
	"github.com/Narutchai01/Project_S-BE/utils"
)

type FavoriteUseCase interface {
	FavoriteCommentThread(thread_id uint, token string) (entities.FavoriteCommentThread, error)
	FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error)
	FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error)
	FavoriteCommnetReviewSkincare(comment_id uint, token string) (entities.FavoriteCommentReviewSkincare, error)
	Favorite(favorite entities.Favorite, type_community string, token string) (entities.Favorite, error)
}

type favoriteService struct {
	repo          repositories.FavoriteRepository
	userRepo      repositories.UserRepository
	threadRepo    repositories.ThreadRepository
	reviewRepo    repositories.ReviewRepository
	commentRepo   repositories.CommentRepository
	communityRepo repositories.CommunityRepository
}

func NewFavoriteUseCase(repo repositories.FavoriteRepository, userRepo repositories.UserRepository, threadRepo repositories.ThreadRepository, reviewRepo repositories.ReviewRepository, commemntRepo repositories.CommentRepository, communityrepo repositories.CommunityRepository) FavoriteUseCase {
	return &favoriteService{repo, userRepo, threadRepo, reviewRepo, commemntRepo, communityrepo}
}

func (service *favoriteService) FavoriteCommentThread(comment_id uint, token string) (entities.FavoriteCommentThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteCommentThread{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteCommentThread{}, errors.New("user not found")
	}

	comment, err := service.commentRepo.GetCommentThread(comment_id)
	if err != nil {
		return entities.FavoriteCommentThread{}, errors.New("comment not found")
	}

	favorite, err := service.repo.FindFavoriteCommentThread(comment.ID, user.ID)
	if err != nil {
		return service.repo.FavoriteCommentThread(comment_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteCommentThread(favorite)

}

func (service *favoriteService) FavoriteThread(thread_id uint, token string) (entities.FavoriteThread, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteThread{}, err
	}
	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteThread{}, errors.New("user not found")
	}

	_, err = service.threadRepo.GetThread(thread_id)
	if err != nil {
		return entities.FavoriteThread{}, errors.New("thread not found")
	}

	favorite, err := service.repo.FindFavoriteThread(thread_id, user.ID)
	if err != nil {
		return service.repo.FavoriteThread(thread_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteThread(favorite)
}

func (service *favoriteService) FavoriteReviewSkincare(review_id uint, token string) (entities.FavoriteReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, errors.New("user not found")
	}

	_, err = service.reviewRepo.GetReviewSkincare(review_id)
	if err != nil {
		return entities.FavoriteReviewSkincare{}, errors.New("review not found")
	}

	favorite, err := service.repo.FindFavoriteReviewSkincare(review_id, user.ID)
	if err != nil {
		return service.repo.FavoriteReviewSkincare(review_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteReviewSkincare(favorite)
}

func (service *favoriteService) FavoriteCommnetReviewSkincare(comment_id uint, token string) (entities.FavoriteCommentReviewSkincare, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, errors.New("user not found")
	}

	comment, err := service.commentRepo.GetCommentThread(comment_id)
	if err != nil {
		return entities.FavoriteCommentReviewSkincare{}, errors.New("comment not found")
	}

	favorite, err := service.repo.FindFavoriteCommentReviewSkincare(comment.ID, user.ID)
	if err != nil {
		return service.repo.FavoriteCommentReviewSkincare(comment_id, user.ID)
	}

	favorite.Status = !favorite.Status
	return service.repo.UpdateFavoriteCommentReviewSkincare(favorite)
}

func (service *favoriteService) Favorite(favorite entities.Favorite, type_community string, token string) (entities.Favorite, error) {

	user_id, err := utils.ExtractToken(token)
	if err != nil {
		return entities.Favorite{}, err
	}

	user, err := service.userRepo.GetUser(user_id)
	if err != nil {
		return entities.Favorite{}, errors.New("user not found")
	}

	favorite.UserID = user.ID

	if favorite.CommunityID != 0 && favorite.CommentID == 0 {
		community_type, err := service.communityRepo.GetCommunityType(strings.ToLower(type_community))
		if err != nil {
			return entities.Favorite{}, errors.New("community not found")
		}

		community, err := service.communityRepo.GetCommunity(favorite.CommunityID, uint64(community_type.ID))
		if err != nil {
			return entities.Favorite{}, errors.New("community not found")
		}

		result, fav_id, err := service.repo.FindFavorite(community.ID, "community_id", uint(user.ID))
		if err != nil {
			return entities.Favorite{}, err
		}

		if !result {
			return service.repo.Favorite(favorite)
		} else {
			return service.repo.DeleteFavorite(fav_id)
		}
	} else if favorite.CommunityID == 0 && favorite.CommentID != 0 {
		comment, err := service.commentRepo.GetComment(favorite.CommentID)
		if err != nil {
			return entities.Favorite{}, errors.New("comment not found")
		}

		result, fav_id, err := service.repo.FindFavorite(comment.ID, "comment_id", user.ID)
		if err != nil {
			return entities.Favorite{}, err
		}

		if !result {
			return service.repo.Favorite(favorite)
		} else {
			return service.repo.DeleteFavorite(fav_id)
		}
	}

	return favorite, nil
}
