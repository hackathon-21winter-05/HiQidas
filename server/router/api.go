package router

import (
	"github.com/hackathon-21winter-05/HiQidas/server/router/heya"
	"github.com/hackathon-21winter-05/HiQidas/server/router/middleware"
	"github.com/hackathon-21winter-05/HiQidas/server/router/oauth"
	"github.com/hackathon-21winter-05/HiQidas/server/router/user"
	"github.com/hackathon-21winter-05/HiQidas/server/router/ws"
)

type APIHandler struct {
	middleware.IMIddleware
	heya.HeyaHandler
	user.UserHandler
	oauth.OauthHandler
	ws.WSHandler
}

func NewAPIHandler(IMIddleware middleware.IMIddleware, heyaHandler heya.HeyaHandler, userHandler user.UserHandler, oauthHandler oauth.OauthHandler, WSHandler ws.WSHandler) *APIHandler {
	return &APIHandler{IMIddleware: IMIddleware, HeyaHandler: heyaHandler, UserHandler: userHandler, OauthHandler: oauthHandler, WSHandler: WSHandler}
}
