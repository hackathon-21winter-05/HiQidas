package streamer

import (
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/heya"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type Streamer struct {
	hs *heya.HeyaStreamer
}

func NewStreamer(ser *service.Service) *Streamer {
	return &Streamer{
		hs: heya.NewHeyaStreamer(ser),
	}
}

func (s *Streamer) Run() {
	go s.hs.Listen()
}

func (s *Streamer) ConnectHeyaWS(c echo.Context) error {
	return s.hs.ConnectHeyaWS(c)
}
