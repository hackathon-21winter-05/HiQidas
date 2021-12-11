package streamer

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/parser"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer/yjs"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
)

type Streamer struct {
	hs *heya.HeyaStreamer
	ys *yjs.YjsStreamer
	ps *parser.ParserStreamer
}

func NewStreamer(c *config.Config, ser *service.Service) *Streamer {
	ps := parser.NewParserStreamer(c, ser)

	return &Streamer{
		hs: heya.NewHeyaStreamer(ser),
		ys: yjs.NewYjsStreamer(ser, ps),
		ps: ps,
	}
}

func (s *Streamer) Run() {
	go s.hs.Listen()
	go s.ys.Listen()
	go s.ps.Listen()
}

func (s *Streamer) ConnectHeyaWS(c echo.Context) error {
	return s.hs.ConnectHeyaWS(c)
}

func (s *Streamer) ConnectYjsWS(c echo.Context) error {
	return s.ys.ConnectYjsWS(c)
}

func (s *Streamer) ConnectParserWS(c echo.Context) error {
	return s.ps.ConnectParserWS(c)
}
