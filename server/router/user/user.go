package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type UserHandlerGroup struct {
	us service.Service
}

func NewUserHandlerGroup(us service.Service) *UserHandlerGroup {
	return &UserHandlerGroup{us: us}
}
