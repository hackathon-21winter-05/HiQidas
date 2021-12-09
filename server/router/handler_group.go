package router

import (
	"net/http"

	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/labstack/echo/v4"
	"github.com/sapphi-red/go-traq"
)

type HandlerGroup interface {
	Setup(*echo.Group)
	Path() string
}

func newHandlerGroups(c *config.Config, ser *service.Service, s *streamer.Streamer) []HandlerGroup {
	cli := traq.NewAPIClient(traq.NewConfiguration())

	hgs := []HandlerGroup{
		user.NewUserHandlerGroup(ser),
		oauth.NewOauthHandlerGroup(c, cli),
		ws.NewWSHandlerGroup(s),
	}

	return hgs
}

// ルーターのハンドラを設定
func (r *Router) setHandlers() {
	api := r.e.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	for _, hg := range r.hgs {
		group := api.Group(hg.Path())
		hg.Setup(group)
	}
}
