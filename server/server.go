package server

import (
	"log"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
)

// サーバー
type Server struct {
	r *router.Router
}

// 新たなサーバーを取得
func NewServer(c *config.Config) *Server {
	server := &Server{
		r: router.NewRouter(c),
	}

	return server
}

// サーバーを起動
func (s *Server) Run() {
	log.Panic(s.r.Run())
}
