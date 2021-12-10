package server

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
)

// サーバー
type Server struct {
	r *router.Router
	s *streamer.Streamer
}

func NewServer(r *router.Router, s *streamer.Streamer) *Server {
	return &Server{r: r, s: s}
}


// サーバーを起動
func (s *Server) Run(c *config.Config) {
	go s.s.Run()
	s.r.Run()
}
