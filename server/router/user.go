package router

import (
	"net/http"

	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/labstack/echo/v4"
)

func (r *Router) GetUsersHandler(c echo.Context) error {
	userIDs, err := r.ser.GetUsersID()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	userStringIDs := uuidsToStrings(userIDs)

	res := &rest.GetUsersResponse{
		UserId: userStringIDs,
	}

	return sendProtobuf(c, http.StatusOK, res)
}
