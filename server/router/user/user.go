package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type UserHandlerGroup struct {
	ser *service.Service
}

func NewUserHandlerGroup(ser *service.Service) *UserHandlerGroup {
	return &UserHandlerGroup{
		ser: ser,
	}
}

func (uh *UserHandlerGroup) Path() string {
	return "/users"
}

func (uh *UserHandlerGroup) Setup(usersApi *echo.Group) {
	usersApi.GET("", uh.GetUsersHandler)
	usersApi.POST("", uh.PostUsersHandler)
}
