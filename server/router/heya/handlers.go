package heya

import (
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
	return nil
}

func (h *HeyaHandleGroup) PostHeyasHandler(c echo.Context) error {
	heyaRequest := rest.PostHeyasRequest{}

	if err:=utils.BindProtobuf(c,&heyaRequest); err!=nil {
		c.Logger().Info(err)
		return echo.NewHTTPError(http.StatusBadRequest,err)
	}

	//service層にデータを加工してもらう
	heyaResponse := rest.PostHeyasResponse{}
	heya,err := h.hs.SaveHeya(heyaRequest.Title,heyaRequest.Description)

	return utils.SendProtobuf(c,http.StatusCreated,&heyaResponse)
}

func (h *HeyaHandleGroup) PutHeyasByIDHandler(c echo.Context) error {
	return nil
}
