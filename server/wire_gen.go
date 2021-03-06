// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InjectServer(c *config.Config) (*Server, error) {
	repositoryRepository, err := repository.NewGormRepository(c)
	if err != nil {
		return nil, err
	}
	serviceService := service.NewService(repositoryRepository)
	middlewareMiddleware := middleware.NewMiddleware(serviceService)
	heyaHandlerGroup := heya.NewHeyaHandleGroup(serviceService)
	userHandlerGroup := user.NewUserHandlerGroup(serviceService)
	configuration := traq.NewConfiguration()
	apiClient := traq.NewAPIClient(configuration)
	oauthHandlerGroup := oauth.NewOauthHandlerGroup(c, serviceService, apiClient)
	streamerStreamer := streamer.NewStreamer(c, serviceService)
	wsHandlerGroup := ws.NewWSHandlerGroup(streamerStreamer)
	apiHandler := router.NewAPIHandler(middlewareMiddleware, heyaHandlerGroup, userHandlerGroup, oauthHandlerGroup, wsHandlerGroup)
	routerRouter := router.NewRouter(apiHandler)
	server := NewServer(routerRouter, streamerStreamer)
	return server, nil
}

// wire.go:

var SuperSet = wire.NewSet(repository.NewGormRepository, wire.Struct(new(repository.GormRepository), "*"), service.NewService, user.NewUserHandlerGroup, heya.NewHeyaHandleGroup, oauth.NewOauthHandlerGroup, wire.NewSet(traq.NewAPIClient, traq.NewConfiguration), middleware.NewMiddleware, ws.NewWSHandlerGroup, streamer.NewStreamer, router.NewAPIHandler, router.NewRouter, NewServer)
