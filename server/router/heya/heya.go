package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type HeyaHandleGroup struct {
	hs *service.Service
}

func NewHeyaHandleGroup(hs *service.Service) *HeyaHandleGroup {
	return &HeyaHandleGroup{hs: hs}
}

type HeyaHandler interface {
	GetHeyasHandler(c echo.Context) error
	GetHeyaHandler(c echo.Context) error
	GetUsersByHeyaIDHandler(c echo.Context) error
	DeleteHeyasByIDHandler(c echo.Context) error
	PostHeyasHandler(c echo.Context) error
	PutHeyasByIDHandler(c echo.Context) error
}