package oauth

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
	"github.com/hackathon-21winter-05/HiQidas/service"
	"github.com/sapphi-red/go-traq"
)

type OauthHandlerGroup struct {
	c   *config.Config
	ser *service.Service
	cli *traq.APIClient
}

func NewOauthHandlerGroup(c *config.Config, ser *service.Service, cli *traq.APIClient) *OauthHandlerGroup {
	return &OauthHandlerGroup{
		c:   c,
		ser: ser,
		cli: cli,
	}
}
