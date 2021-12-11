package heya

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
)

// GetHeyasHandler GET /heyas
func (h *HeyaHandlerGroup) GetHeyasHandler(c echo.Context) error {
	heyas, err := h.hs.GetHeyas(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	var resHeya []*rest.Heya
	for _, heya := range heyas {
		resHeya = append(resHeya, &rest.Heya{
			Id:           heya.ID.String(),
			Title:        heya.Title,
			Description:  heya.Description,
			CreatorId:    heya.CreatorID.String(),
			CreatorName:  heya.Creator.Name,
			LastEditorId: heya.LastEditorID.String(),
			CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
			UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
		})
	}
	res := rest.GetHeyasResponse{
		Heyas: &rest.Heyas{Heyas: resHeya},
	}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetHeyaHandler GET /heyas/:heyaID
func (h *HeyaHandlerGroup) GetHeyaHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	heya, err := h.hs.GetHeyaByID(c.Request().Context(), heyaUUID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	resHeya := rest.Heya{
		Id:           heya.ID.String(),
		Title:        heya.Title,
		Description:  heya.Description,
		CreatorId:    heya.CreatorID.String(),
		LastEditorId: heya.LastEditorID.String(),
		CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
		UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
	}
	res := rest.GetHeyasHeyaIdResponse{Heya: &resHeya}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetUsersByHeyaIDHandler GET /heyas/:heyaID/users
func (h *HeyaHandlerGroup) GetUsersByHeyaIDHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	clients := h.hs.GetHeyaClientsIDByHeyaID(heyaUUID)
	userIDs := make([]uuid.UUID, 0)
	for _, client := range clients {
		userIDs = append(userIDs, client.UserID)
	}

	res := rest.GetHeyasHeyaIdUsersResponse{UserId: utils.UuidsToStrings(userIDs)}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// DeleteHeyasByIDHandler DELETE /heyas/:heyaID
func (h *HeyaHandlerGroup) DeleteHeyasByIDHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err = h.hs.DeleteHeya(c.Request().Context(), heyaUUID); err != nil {
		if errors.Is(err, repository.ErrNoRecordDeleted) {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

// PostHeyasHandler POST /heyas
func (h *HeyaHandlerGroup) PostHeyasHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	heyaRequest := rest.PostHeyasRequest{}

	if err := utils.BindProtobuf(c, &heyaRequest); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	heya, err := h.hs.CreateHeya(c.Request().Context(), userID, heyaRequest.Title, heyaRequest.Description)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	resHeya := rest.Heya{
		Id:           heya.ID.String(),
		Title:        heya.Title,
		Description:  heya.Description,
		CreatorId:    heya.CreatorID.String(),
		LastEditorId: heya.LastEditorID.String(),
		CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
		UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
	}
	res := rest.PostHeyasResponse{
		Heya: &resHeya,
	}
	return utils.SendProtobuf(c, http.StatusCreated, &res)
}

// PutHeyasByIDHandler PUT /heyas/:heyaID
func (h *HeyaHandlerGroup) PutHeyasByIDHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	heyaRequest := rest.PutHeyasHeyaIdRequest{}

	if err := utils.BindProtobuf(c, &heyaRequest); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	heya := model.NullHeya{
		Title:       sql.NullString{String: heyaRequest.Title, Valid: true},
		Description: sql.NullString{String: heyaRequest.Description, Valid: true},
	}

	if err = h.hs.PutHeyaByID(c.Request().Context(), &heya, heyaUUID, userID); err != nil {
		if errors.Is(err, repository.ErrNoRecordUpdated) {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

// PutFavoriteByHeyaIDHandler PUT /heyas/{heyaID}/favorite
func (h *HeyaHandlerGroup) PutFavoriteByHeyaIDHandler(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	userIDstr := sess.Values["userid"].(string)
	userID, err := uuid.FromString(userIDstr)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	heyaID := c.Param("heyaID")

	uuidHeyaID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := rest.PutHeyasUserIdFavoriteRequest{}

	if err = utils.BindProtobuf(c, &req); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to bind ", err)
	}

	if err = h.hs.PutFavoriteByHeyaID(c.Request().Context(), uuidHeyaID, userID, req.IsFavorite); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
