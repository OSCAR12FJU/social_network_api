package service

import (
	"api_red_social/model"
	"api_red_social/repository"
	"errors"
)

type FollowService struct {
	Repo *repository.FollowRepository
}

func NewFollowService(followRepo *repository.FollowRepository) *FollowService {
	return &FollowService{
		Repo: followRepo,
	}
}

func (service *FollowService) SaveFollow(userID, followedID string) (*model.Follow, error) {
	if userID == "" || followedID == "" {
		return nil, errors.New("faltan datos para seguir al usuario")
	}
	follow := &model.Follow{
		UserID:   userID,
		FollowID: followedID,
	}

	return service.Repo.SaveFollow(follow)
}

func (service *FollowService) DeleteFollow(userID, followedID string) error {
	err := service.Repo.DeleteFollow(userID, followedID)
	if err != nil {
		return errors.New("no has dejado de seguis a nadie")
	}
	return nil
}

func (service *FollowService) GetFollowing(userID string, page int, itemsPerPage int) ([]model.Follow, int, error) {
	follows, total, err := service.Repo.FindFollowing(userID, page, itemsPerPage)
	if err != nil {
		return nil, 0, err
	}
	return follows, total, nil
}

func (service *FollowService) FollowUserIds(userID string) ([]string, []string, error) {
	// Implement your logic to get the user IDs that the user is following and that follow the user
	return nil, nil, nil
}
