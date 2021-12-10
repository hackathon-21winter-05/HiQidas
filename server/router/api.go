package router

import (
	"github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
)

type APIHandler struct {
	*middleware.Middleware
	*heya.HeyaHandlerGroup
	*user.UserHandlerGroup
	*oauth.OauthHandlerGroup
	*ws.WSHandlerGroup
}

func NewAPIHandler(
	Middleware *middleware.Middleware,
	heyaHandlerGroup *heya.HeyaHandlerGroup,
	userHandlerGroup *user.UserHandlerGroup,
	oauthHandlerGroup *oauth.OauthHandlerGroup,
	WSHandlerGroup *ws.WSHandlerGroup,
) *APIHandler {
	return &APIHandler{
		Middleware:        Middleware,
		HeyaHandlerGroup:  heyaHandlerGroup,
		UserHandlerGroup:  userHandlerGroup,
		OauthHandlerGroup: oauthHandlerGroup,
		WSHandlerGroup:    WSHandlerGroup,
	}
}
