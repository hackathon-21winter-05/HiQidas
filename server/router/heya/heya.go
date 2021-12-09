package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/service/heya"
	"github.com/labstack/echo/v4"
)

type HeyaHandleGroup struct {
	hs heya.HeyaService
}

func NewHeyaHandleGroup(hs heya.HeyaService) *HeyaHandleGroup {
	return &HeyaHandleGroup{hs: hs}
}

func (h *HeyaHandleGroup) Setup(heyasApi *echo.Group) {
	heyasApi.GET("",h.GetHeyasHandler)
}