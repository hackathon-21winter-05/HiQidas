package user

import (
	"net/http"

	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
)

func (uh *UserHandlerGroup) GetUsersHandler(c echo.Context) error {
	userIDs, err := uh.s.GetUsersID(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	userStringIDs := utils.UuidsToStrings(userIDs)

	res := &rest.GetUsersResponse{
		UserId: userStringIDs,
	}

	return utils.SendProtobuf(c, http.StatusOK, res)
}
