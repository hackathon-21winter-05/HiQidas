package oauth

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/labstack/echo/v4"
	"github.com/sapphi-red/go-traq"
)

type OauthHandlerGroup struct {
	c   *config.Config
	cli *traq.APIClient
}

func NewOauthHandlerGroup(c *config.Config, cli *traq.APIClient) *OauthHandlerGroup {
	return &OauthHandlerGroup{
		c:   c,
		cli: cli,
	}
}

func (oh *OauthHandlerGroup) Path() string {
	return "/oauth"
}

func (oh *OauthHandlerGroup) Setup(oauthApi *echo.Group) {
	oauthApi.GET("/callback", oh.GetOauthCallbackHandler)
	oauthApi.POST("/code", oh.PostOauthCodeHandler)
}
