package server

import (
	"log"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/rest"
	"github.com/labstack/echo/v4"
)

// サーバー
type Server struct {
	e *echo.Echo
	h rest.RestHandlers
}

// 新たなサーバーを取得
func NewServer(c *config.Config) *Server {
	server := &Server{
		e: newEcho(),
		h: rest.NewRestHandlers(c),
	}

	server.addHandlers()

	return server
}

// サーバーを起動
func (s *Server) Run() {
	log.Panic(s.e.Start(":7070"))
}
