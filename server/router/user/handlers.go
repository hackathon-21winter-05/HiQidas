package user

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo-contrib/session"
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
	sess, err := session.Get("session", c)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get session", err)
	}

	accessToken := sess.Values["accessToken"]
	if accessToken == nil {
		//こんな感じ？
	}

	UserID := sess.Values["userID"].(uuid.UUID)

	user, err := uh.s.GetUserByID(c.Request().Context(), UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := rest.User{
		Id:   UserID.String(),
		Name: user.Name,
	}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetHeyasByMeHandler GET /users/me/heyas
func (uh *UserHandlerGroup) GetHeyasByMeHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get session", err)
	}

	accessToken := sess.Values["accessToken"]
	if accessToken == nil {
		//こんな感じ？
	}

	UserID := sess.Values["userID"].(uuid.UUID)

	heyas, err := uh.s.GetHeyaByUserMe(c.Request().Context(), UserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var rheyas []*rest.Heya

	for _, heya := range heyas {
		rheyas = append(rheyas, &rest.Heya{
			Id:           heya.ID.String(),
			Title:        heya.Title,
			Description:  heya.Description,
			CreatorId:    heya.CreatorID.String(),
			LastEditorId: heya.LastEditorID.String(),
			CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
			UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
		})
	}
	res := rest.GetHeyasResponse{
		Heyas: &rest.Heyas{Heyas: rheyas},
	}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetFavoriteUsersMeHandler GET /users/me/favorites
func (uh *UserHandlerGroup) GetFavoriteUsersMeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "favorite")
}

// GetUsersByIDHandler GET /users/{userID}
func (uh *UserHandlerGroup) GetUsersByIDHandler(c echo.Context) error {
	userID := c.Param("userID")
	sess, err := session.Get("session", c)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get session", err)
	}

	accessToken := sess.Values["accessToken"]
	if accessToken == nil {
		//こんな感じ？
	}

	uuidUserID, err := uuid.FromString(userID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user, err := uh.s.GetUserByID(c.Request().Context(), uuidUserID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := rest.User{
		Id:   userID,
		Name: user.Name,
	}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// PostUsersHandler POST /users
func (uh *UserHandlerGroup) PostUsersHandler(c echo.Context) error {
	req := rest.PostUsersResponse{}

	if err := utils.BindProtobuf(c, &req); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}

	res,err := uh.s.CreateUser(c.Request().Context(),req.User.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	protoRes := rest.PostUsersResponse{User: &rest.User{
		Id:   res.ID.String(),
		Name: res.Name,
	}}

	return utils.SendProtobuf(c,http.StatusCreated,&protoRes)
}
