package users

import "api_red_social/internal/ports"

type Handler struct {
	UserService ports.UsersService
}
