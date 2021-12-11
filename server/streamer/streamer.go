package streamer

import (
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/yjs"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type Streamer struct {
	hs *heya.HeyaStreamer
	ys *yjs.YjsStreamer
}

func NewStreamer(ser *service.Service) *Streamer {
	return &Streamer{
		hs: heya.NewHeyaStreamer(ser),
		ys: yjs.NewYjsStreamer(ser),
	}
}

func (s *Streamer) Run() {
	go s.hs.Listen()
	go s.ys.Listen()
}

func (s *Streamer) ConnectHeyaWS(c echo.Context) error {
	return s.hs.ConnectHeyaWS(c)
}

func (s *Streamer) ConnectYjsWS(c echo.Context) error {
	return s.ys.ConnectYjsWS(c)
}
