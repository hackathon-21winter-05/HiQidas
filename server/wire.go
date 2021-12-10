//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
	"github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/sapphi-red/go-traq"
)

var SuperSet = wire.NewSet(
	repository.NewGormRepository,
	wire.Struct(new(repository.GormRepository), "*"),
	service.NewService,

	user.NewUserHandlerGroup,
	wire.Bind(new(user.UserHandler),new(*user.UserHandlerGroup)),
	heya.NewHeyaHandleGroup,
	wire.Bind(new(heya.HeyaHandler),new(*heya.HeyaHandleGroup)),
	oauth.NewOauthHandlerGroup,
	wire.Bind(new(oauth.OauthHandler),new(*oauth.OauthHandlerGroup)),
	wire.NewSet(traq.NewAPIClient, traq.NewConfiguration),
	middleware.NewMiddleware,
	wire.Bind(new(middleware.IMIddleware),new(*middleware.Middleware)),
	ws.NewWSHandlerGroup,
	wire.Bind(new(ws.WSHandler),new(*ws.WSHandlerGroup)),
	streamer.NewStreamer,
	router.NewAPIHandler,
	router.NewRouter,
	NewServer,
)

func InjectServer(c *config.Config) (*Server, error) {
	wire.Build(SuperSet)
	return nil, nil
}
