package router

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sapphi-red/go-traq"
)

// ルーター
type Router struct {
	e   *echo.Echo
	c   *config.Config
	cli *traq.APIClient
}

// 新しいルーターを生成
func NewRouter(c *config.Config, s *streamer.Streamer) *Router {
	r := &Router{
		e:   newEcho(),
		c:   c,
		cli: traq.NewAPIClient(traq.NewConfiguration()),
	}

	r.setHandlers(s)

	return r
}

// 設定済みの新しいEchoインスタンスを生成
func newEcho() *echo.Echo {
	e := echo.New()

	// ログの設定
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n"}))

	// セッションの設定
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	return e
}

// ルーターのハンドラを設定
func (r *Router) setHandlers(s *streamer.Streamer) {
	api := r.e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})

		oauthApi := api.Group("/oauth")
		{
			oauthApi.GET("/callback", r.GetOauthCallbackHandler)
			oauthApi.POST("/code", r.PostOauthCodeHandler)
		}

		wsApi := api.Group("/ws")
		{
			wsApi.GET("/{roomid}", s.ConnectWs)
		}
	}
}

// ルーターを起動
func (r *Router) Run() error {
	return r.e.Start(":7070")
}
