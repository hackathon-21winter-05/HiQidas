package server

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

// サーバー
type Server struct {
	r *router.Router
	s *streamer.Streamer
}

// 新たなサーバーを取得
func NewServer(c *config.Config) *Server {
	server, err := injectServer(c)
	if err != nil {
		log.Panic(err)
	}
	e := server.r.NewEcho()

	echoApi := e.Group("/api")
	{
		echoApi.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		userApi := echoApi.Group("/users")
		{
			userApi.GET("", server.r.Api.GetUsersHandler)
		}

		heyaApi := echoApi.Group("/heyas")
		{
			heyaApi.GET("", server.r.Api.GetHeyasHandler)
			heyaApi.GET("/:heyaID", server.r.Api.GetHeyaHandler)
			heyaApi.GET("/:heyaID/users", server.r.Api.GetUsersByHeyaIDHandler)
			heyaApi.POST("/", server.r.Api.PostHeyasHandler)
			heyaApi.DELETE("/:heyaID", server.r.Api.DeleteHeyasByIDHandler)
			heyaApi.PUT("/:heyaID", server.r.Api.PutHeyasByIDHandler)
		}

		oauthApi := echoApi.Group("/oauth")
		{
			oauthApi.GET("/callback", server.r.Api.GetOauthCallbackHandler)
			oauthApi.POST("/code", server.r.Api.PostOauthCodeHandler)
		}

		echoApi.GET("*", func(c echo.Context) error {
			return c.String(http.StatusNotImplemented, "Not Implemented")
		})
	}

	return server
}

// サーバーを起動
func (s *Server) Run() {
	s.s.Run()
	err := s.r.Run()
	if err != nil {
		log.Panic(err)
	}
}
