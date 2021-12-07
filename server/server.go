package server

import (
	"log"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
)

// サーバー
type Server struct {
	r *router.Router
	s *streamer.Streamer
}

// 新たなサーバーを取得
func NewServer(c *config.Config) *Server {
	s := streamer.NewStreamer()
	r := router.NewRouter(c, s)

	server := &Server{
		r: r,
		s: s,
	}

	return server
}

// サーバーを起動
func (s *Server) Run() {
	s.s.Run()
	log.Panic(s.r.Run())
}
