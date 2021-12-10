package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type HeyaHandlerGroup struct {
	hs *service.Service
}

func NewHeyaHandleGroup(hs *service.Service) *HeyaHandlerGroup {
	return &HeyaHandlerGroup{hs: hs}
}
