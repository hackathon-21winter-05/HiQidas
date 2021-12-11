package user

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetUsersHandler GET /users
func (uh *UserHandlerGroup) GetUsersHandler(c echo.Context) error {
	userIDs, err := uh.s.GetUsersID(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
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
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user, err := uh.s.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := &rest.GetUsersMeResponse{
		Me: &rest.User{
			Id:         userID.String(),
			Name:       user.Name,
			IconFileId: user.IconFileID.UUID.String(),
		}}

	return utils.SendProtobuf(c, http.StatusOK, res)
}

// GetHeyasByMeHandler GET /users/me/heyas
func (uh *UserHandlerGroup) GetHeyasByMeHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	heyas, err := uh.s.GetHeyaByUserMe(c.Request().Context(), userID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var resHeyas []*rest.Heya

	for _, heya := range heyas {
		resHeyas = append(resHeyas, &rest.Heya{
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
		Heyas: &rest.Heyas{Heyas: resHeyas},
	}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetFavoriteUsersMeHandler GET /users/me/favorites
func (uh *UserHandlerGroup) GetFavoriteUsersMeHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusInternalServerError,err)
	}

	favorites, err := uh.s.GetUserMeFavorites(c.Request().Context(), userID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	favoIDs := []string{}
	for _, favorite := range favorites {
		favoIDs = append(favoIDs, favorite.HeyaID.String())
	}
	res := rest.GetUsersMeFavoritesRequest{FavoriteHeyaId: favoIDs}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetUsersByIDHandler GET /users/{userID}
func (uh *UserHandlerGroup) GetUsersByIDHandler(c echo.Context) error {
	userID := c.Param("userID")

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

	res := &rest.GetUsersUserIdResponse{
		User: &rest.User{
			Id:         userID,
			Name:       user.Name,
			IconFileId: user.IconFileID.UUID.String(),
		}}

	return utils.SendProtobuf(c, http.StatusOK, res)
}

// PostUsersHandler POST /users
func (uh *UserHandlerGroup) PostUsersHandler(c echo.Context) error {
	req := rest.PostUsersRequest{}

	if err := utils.BindProtobuf(c, &req); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := uh.s.CreateUser(c.Request().Context(), req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sess.Values["userid"] = res.ID.String()
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	protoRes := rest.PostUsersResponse{User: &rest.User{
		Id:   res.ID.String(),
		Name: res.Name,
	}}

	return utils.SendProtobuf(c, http.StatusCreated, &protoRes)
}
