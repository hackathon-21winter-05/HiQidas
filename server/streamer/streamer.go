package streamer

import (
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/heya"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type Streamer struct {
	*heya.HeyaStreamer
}

func NewStreamer(ser *service.Service) *Streamer {
	return &Streamer{
		HeyaStreamer: heya.NewHeyaStreamer(ser),
	}
}

func (s *Streamer) Run() {
	s.HeyaStreamer.Listen()
}

func (s *Streamer) ConnectHeyaWS(c echo.Context) error {
	return s.HeyaStreamer.ConnectHeyaWS(c)
}
