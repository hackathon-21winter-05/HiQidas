package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type UserHandlerGroup struct {
	s *service.Service
}

func NewUserHandlerGroup(s *service.Service) *UserHandlerGroup {
	return &UserHandlerGroup{s: s}
}
