// +build wireinject
package router

import (
	"github.com/google/wire"
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	heya2 "github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	user2 "github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
	"github.com/hackathon-21winter-05/HiQidas/server/streamer"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/hackathon-21winter-05/HiQidas/service/heya"
)


var SuperSet = wire.NewSet(
	repository.NewGormRepository,
	wire.Struct(new(repository.GormRepository),"*"),

	heya.NewHeyaServiceImpl,
	wire.Bind(new(heya.HeyaService),new(*heya.HeyaServiceImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService),new(*service.UserServiceImpl)),
	NewAPI,
	heya2.NewHeyaHandleGroup,
	user2.NewUserHandlerGroup,
	ws.NewWSHandlerGroup,
	streamer.NewStreamer,
	middleware.NewMiddleware,
	)
func InjectAPIServer(c *config.Config) (*API,error) {
	wire.Build(SuperSet)

	return nil,nil
}