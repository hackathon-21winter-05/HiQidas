package router

import (
	"github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
)

type API struct {
	*middleware.Middleware
	*heya.HeyaHandleGroup
	*user.UserHandlerGroup
	*oauth.OauthHandlerGroup
	*ws.WSHandlerGroup
}

func NewAPI(middleware *middleware.Middleware, heyaHandleGroup *heya.HeyaHandleGroup, userHandlerGroup *user.UserHandlerGroup, oauthHandlerGroup *oauth.OauthHandlerGroup, WSHandlerGroup *ws.WSHandlerGroup) *API {
	return &API{Middleware: middleware, HeyaHandleGroup: heyaHandleGroup, UserHandlerGroup: userHandlerGroup, OauthHandlerGroup: oauthHandlerGroup, WSHandlerGroup: WSHandlerGroup}
}

