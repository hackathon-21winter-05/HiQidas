package router

import (
	"github.com/gorilla/sessions"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

// ルーター
type Router struct {
	e *echo.Echo
}

// 新しいルーターを生成
func NewRouter(c *config.Config) *Router {
	r := &Router{
		e: newEcho(),
	}
	_api, err := InjectAPIHandlers(c)
	if err != nil {
		log.Error(err)
	}
	api := *_api

	echoApi := r.e.Group("/api")
	{
		echoApi.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		userApi := echoApi.Group("/users")
		{
			userApi.GET("", api.GetUsersHandler)
		}

		heyaApi := echoApi.Group("/heyas")
		{
			heyaApi.GET("", api.GetHeyasHandler)
			heyaApi.GET("/:heyaID", api.GetHeyaHandler)
			heyaApi.GET("/:heyaID/users", api.GetUsersByHeyaIDHandler)
			heyaApi.POST("/", api.PostHeyasHandler)
			heyaApi.DELETE("/:heyaID", api.DeleteHeyasByIDHandler)
			heyaApi.PUT("/:heyaID", api.PutHeyasByIDHandler)
		}

		oauthApi := echoApi.Group("/oauth")
		{
			oauthApi.GET("/callback", api.GetOauthCallbackHandler)
			oauthApi.POST("/code", api.PostOauthCodeHandler)
		}

		echoApi.GET("*", func(c echo.Context) error {
			return c.String(http.StatusNotImplemented, "Not Implemented")
		})
	}

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
