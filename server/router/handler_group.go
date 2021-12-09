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
}

type HandlerGroups struct {
	userHG  HandlerGroup
	oauthHG HandlerGroup
	wsHG    HandlerGroup
}

func newHandlerGroups(c *config.Config, ser *service.Service, s *streamer.Streamer) *HandlerGroups {
	cli := traq.NewAPIClient(traq.NewConfiguration())

	hgs := &HandlerGroups{
		userHG:  user.NewUserHandlerGroup(ser),
		oauthHG: oauth.NewOauthHandlerGroup(c, cli),
		wsHG:    ws.NewWSHandlerGroup(s),
	}

	return hgs
}

// ルーターのハンドラを設定
func (r *Router) setHandlers() {
	api := r.e.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	userApi := api.Group("/user")
	r.hgs.userHG.Setup(userApi)

	oauthApi := api.Group("/oauth")
	r.hgs.oauthHG.Setup(oauthApi)

	wsApi := api.Group("/ws")
	r.hgs.wsHG.Setup(wsApi)
}
