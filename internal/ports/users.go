package ports

import "api_red_social/internal/domain"

type UsersService interface {
	Create(user domain.Users) (id interface{}, err error)
	Profile(user string) (id interface{}, err error)
}

type UsersRepository interface {
	Insert(user domain.Users) (id interface{}, err error)
	ProfileId(userId string) (id interface{}, err error)
}
