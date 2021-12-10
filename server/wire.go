//+build wireinject
package server

import (
	"github.com/google/wire"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/server/router"
	heya2 "github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	user2 "github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/hackathon-21winter-05/HiQidas/service/heya"
	"github.com/sapphi-red/go-traq"
)


var SuperSet = wire.NewSet(
	repository.NewGormRepository,
	wire.Struct(new(repository.GormRepository),"*"),

	heya.NewHeyaServiceImpl,
	wire.Bind(new(heya.HeyaService),new(*heya.HeyaServiceImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService),new(*service.UserServiceImpl)),

	router.NewAPI,
	heya2.NewHeyaHandleGroup,
	user2.NewUserHandlerGroup,
	ws.NewWSHandlerGroup,
	oauth.NewOauthHandlerGroup,
	wire.NewSet(traq.NewAPIClient,traq.NewConfiguration),
	streamer.NewStreamer,
	middleware.NewMiddleware,
	)

func injectAPIHandlers(c *config.Config) (*router.APIHandlers,error) {
	wire.Build(SuperSet)

	return nil,nil
}