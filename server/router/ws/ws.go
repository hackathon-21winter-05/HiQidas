package ws

import (
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/labstack/echo/v4"
)

type WSHandlerGroup struct {
	s *streamer.Streamer
}

func NewWSHandlerGroup(s *streamer.Streamer) *WSHandlerGroup {
	return &WSHandlerGroup{
		s: s,
	}
}

func (wh *WSHandlerGroup) Path() string {
	return "/ws"
}

func (wh *WSHandlerGroup) Setup(wsApi *echo.Group) {
	wsApi.GET("/heya/:heyaid", wh.ConnectHeyaWS)
}
