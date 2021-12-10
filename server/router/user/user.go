package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type UserHandlerGroup struct {
	us service.UserService
}

func NewUserHandlerGroup(us service.UserService) *UserHandlerGroup {
	return &UserHandlerGroup{
		us: us,
	}
}
