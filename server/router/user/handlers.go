package user

import (
	"net/http"

	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
)

func (uh *UserHandlerGroup) GetUsersHandler(c echo.Context) error {
	userIDs, err := uh.ser.GetUsersID()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	userStringIDs := utils.UuidsToStrings(userIDs)

	res := &rest.GetUsersResponse{
		UserId: userStringIDs,
	}

	return utils.SendProtobuf(c, http.StatusOK, res)
}

func (uh *UserHandlerGroup) PostUsersHandler(c echo.Context) error {
	usersData := &rest.PostUsersRequest{}
	err := utils.BindProtobuf(c, usersData)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	created, err := uh.ser.CreateUser(usersData.Name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	res := &rest.PostUsersResponse{
		User: &rest.User{
			Id:   created.ID.String(),
			Name: created.Name,
		},
	}

	return utils.SendProtobuf(c, http.StatusCreated, res)
}
