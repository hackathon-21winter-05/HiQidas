package user

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type UserHandlerGroup struct {
	us service.UserService
}

func NewUserHandlerGroup(us service.UserService) *UserHandlerGroup {
	return &UserHandlerGroup{
		us: us,
	}
}

func (uh *UserHandlerGroup) Path() string {
	return "/users"
}

func (uh *UserHandlerGroup) Setup(usersApi *echo.Group) {
	usersApi.GET("", uh.GetUsersHandler)
}
