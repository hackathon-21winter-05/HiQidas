package router

import (
	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// ルーター
type Router struct {
	e   *echo.Echo
	hgs []HandlerGroup
}

// 新しいルーターを生成
func NewRouter(c *config.Config, s *streamer.Streamer, ser *service.Service) *Router {
	r := &Router{
		e:   newEcho(),
		hgs: newHandlerGroups(c, ser, s),
	}

	r.setHandlers()

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

// ルーターを起動
func (r *Router) Run() error {
	return r.e.Start(":7070")
}
