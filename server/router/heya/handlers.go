package heya

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetHeyasHandler GET /heyas
func (h *HeyaHandleGroup) GetHeyasHandler(c echo.Context) error {
	heyas, err := h.hs.GetHeyas(c.Request().Context())
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

// GetHeyaHandler GET /heyas/:heyaID
func (h *HeyaHandleGroup) GetHeyaHandler(c echo.Context) error {
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
	rheya := rest.Heya{
		Id:           heya.ID.String(),
		Title:        heya.Title,
		Description:  heya.Description,
		CreatorId:    heya.CreatorID.String(),
		LastEditorId: heya.LastEditorID.String(),
		CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
		UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
	}
	res := rest.GetHeyasHeyaIdResponse{Heya: &rheya}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// GetUsersByHeyaIDHandler GET /heyas/:heyaID/users
func (h *HeyaHandleGroup) GetUsersByHeyaIDHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	userIDs, err := h.hs.GetUsersByHeyaID(c.Request().Context(), heyaUUID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := rest.GetHeyasHeyaIdUsersResponse{UserId: utils.UuidsToStrings(userIDs)}

	return utils.SendProtobuf(c, http.StatusOK, &res)
}

// DeleteHeyasByIDHandler DELETE /heyas/:heyaID
func (h *HeyaHandleGroup) DeleteHeyasByIDHandler(c echo.Context) error {
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
func (h *HeyaHandleGroup) PostHeyasHandler(c echo.Context) error {
	heyaRequest := rest.PostHeyasRequest{}

	if err := utils.BindProtobuf(c, &heyaRequest); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	//TODO:セッションor MiddlewareからUserIDをもってこれるようにする
	heya, err := h.hs.CreateHeya(c.Request().Context(), uuid.Nil, heyaRequest.Title, heyaRequest.Description)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	rheya := rest.Heya{
		Id:           heya.ID.String(),
		Title:        heya.Title,
		Description:  heya.Description,
		CreatorId:    heya.CreatorID.String(),
		LastEditorId: heya.LastEditorID.String(),
		CreatedAt:    utils.TimeStampToTIme(heya.CreatedAt),
		UpdatedAt:    utils.TimeStampToTIme(heya.UpdatedAt),
	}
	heyaResponse := rest.PostHeyasResponse{
		Heya: &rheya,
	}
	return utils.SendProtobuf(c, http.StatusCreated, &heyaResponse)
}

// PutHeyasByIDHandler PUT /heyas/:heyaID
func (h *HeyaHandleGroup) PutHeyasByIDHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
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

	//TODO:セッションor MiddlewareからUserIDをもってこれるようにする
	if err = h.hs.PutHeyaByID(c.Request().Context(), &heya, heyaUUID, uuid.Nil); err != nil {
		if errors.Is(err, repository.ErrNoRecordUpdated) {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}