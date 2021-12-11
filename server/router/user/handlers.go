package user

import (
	"net/http"

	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
)

// GetUsersHandler GET /users
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

// GetUsersMeHandler GET /users/me
func (uh *UserHandlerGroup) GetUsersMeHandler(c echo.Context) error {
	panic("implement me")
}

// GetHeyasByMeHandler GET /users/me/heyas
func (uh *UserHandlerGroup) GetHeyasByMeHandler(c echo.Context) error {
	panic("implement me")
}

// GetFavoriteUsersMeHandler GET /users/me/favorites
func (uh *UserHandlerGroup) GetFavoriteUsersMeHandler(c echo.Context) error {
	panic("implement me")
}

// GetUsersByIDHandler GET /users/{userID}
func (uh *UserHandlerGroup) GetUsersByIDHandler(c echo.Context) error {
	panic("implement me")
}

// PostUsersHandler POST /users
func (uh *UserHandlerGroup) PostUsersHandler(c echo.Context) error {
	panic("implement me")
}
