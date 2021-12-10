package oauth

import (
	"github.com/hackathon-21winter-05/HiQidas/config"
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
