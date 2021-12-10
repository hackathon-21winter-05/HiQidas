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

type WSHandler interface {
	ConnectHeyaWS(c echo.Context) error
}
