package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type UserHandlerGroup struct {
	s *service.Service
}

func NewUserHandlerGroup(s *service.Service) *UserHandlerGroup {
	return &UserHandlerGroup{s: s}
}

type UserHandler interface {
	GetUsersHandler(c echo.Context) error
}