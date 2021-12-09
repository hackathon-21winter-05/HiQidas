package heya

import (
	"errors"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/server/protobuf/rest"
	"github.com/hackathon-21winter-05/HiQidas/server/router/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HeyaHandleGroup) GetHeyasHandler(c echo.Context) error {
	return nil
}

func (h *HeyaHandleGroup) GetHeyasByIDHandler(c echo.Context) error {
	return nil
}

func (h *HeyaHandleGroup) GetUsersByHeyaIDHandler(c echo.Context) error {
	return nil
}

func (h *HeyaHandleGroup) DeleteHeyasByIDHandler(c echo.Context) error {
	heyaID := c.Param("heyaID")
	heyaUUID, err := uuid.FromString(heyaID)
	if err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err = h.hs.DeleteHeya(heyaUUID); err != nil {
		if errors.Is(err, model.ErrNoRecordDeleted) {
			c.Logger().Info(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *HeyaHandleGroup) PostHeyasHandler(c echo.Context) error {
	heyaRequest := rest.PostHeyasRequest{}

	if err := utils.BindProtobuf(c, &heyaRequest); err != nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	//TODO:セッションからUserIDをもってこれるようにする
	heya, err := h.hs.CreateHeya(uuid.Nil, heyaRequest.Title, heyaRequest.Description)
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

func (h *HeyaHandleGroup) PutHeyasByIDHandler(c echo.Context) error {
	return nil
}
