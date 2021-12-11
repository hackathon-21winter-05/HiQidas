package router

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// ルーター
type Router struct {
	e   *echo.Echo
	api *APIHandler
}

func NewRouter(api *APIHandler) *Router {
	e := newEcho()

	echoApi := e.Group("/api")
	{
		echoApi.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})

		userApi := echoApi.Group("/users", api.CheckLogin)
		{
			userApi.GET("", api.GetUsersHandler)
			userApi.GET("/:userID", api.GetUsersByIDHandler)
			userApi.GET("/me", api.GetUsersMeHandler)
			userApi.GET("/me/favorites", api.GetFavoriteUsersMeHandler)
			userApi.GET("/me/heyas", api.GetHeyasByMeHandler)
		}

		unAuthUserApi := echoApi.Group("/users")
		{
			unAuthUserApi.POST("", api.PostUsersHandler)
		}

		heyaApi := echoApi.Group("/heyas", api.CheckLogin)
		{
			heyaApi.GET("", api.GetHeyasHandler)
			heyaApi.GET("/:heyaID", api.GetHeyaHandler)
			heyaApi.GET("/:heyaID/users", api.GetUsersByHeyaIDHandler)
			heyaApi.POST("", api.PostHeyasHandler)
			heyaApi.DELETE("/:heyaID", api.DeleteHeyasByIDHandler, api.DeleteHeyaCheckAdmin)
			heyaApi.PUT("/:heyaID", api.PutHeyasByIDHandler)
			heyaApi.PUT("/:heyaID/favorite", api.PutFavoriteByHeyaIDHandler)
		}

		oauthApi := echoApi.Group("/oauth")
		{
			oauthApi.GET("/callback", api.GetOauthCallbackHandler)
			oauthApi.POST("/code", api.PostOauthCodeHandler)
		}

		wsApi := echoApi.Group("/ws", api.CheckLogin)
		{
			wsApi.GET("/heya/:heyaid", api.ConnectHeyaWS)
			wsApi.GET("/yjs/:hiqidashiid", api.ConnectYjsWS)
			wsApi.GET("/parser", api.ConnectParserWS)
		}

		echoApi.GET("*", func(c echo.Context) error {
			return c.String(http.StatusNotImplemented, "Not Implemented")
		})

	}
	return &Router{e: e, api: api}
}

func (r *Router) Run() {
	r.e.Logger.Fatal(r.e.Start(":7070"))
}

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
