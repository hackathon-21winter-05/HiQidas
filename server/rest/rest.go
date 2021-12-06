package rest

import "github.com/hackathon-21winter-05/HiQidas/config"

// REST APIハンドラー群
type RestHandlers interface {
	oauthHandlers
}

// REST APIハンドラー群 実装
type restHandlersImpl struct {
	c *config.Config
}

// 新しいREST APIハンドラー群を生成
func NewRestHandlers(c *config.Config) RestHandlers {
	return &restHandlersImpl{
		c: c,
	}
}
