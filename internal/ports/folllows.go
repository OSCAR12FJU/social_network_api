package ports

import "api_red_social/internal/domain"

type FollowService interface {
	Create(follow domain.Follows) (id interface{}, err error)
}

type FollowRepository interface {
	Insert(follow domain.Follows) (id interface{}, err error)
}
