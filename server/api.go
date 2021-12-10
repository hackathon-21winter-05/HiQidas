package server

import (
	"github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
)

type API struct {
	*middleware.Middleware
	*heya.HeyaHandleGroup
	*user.UserHandlerGroup
}
