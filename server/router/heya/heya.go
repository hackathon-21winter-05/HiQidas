package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/service"
)

type HeyaHandleGroup struct {
	hs *service.Service
}

func NewHeyaHandleGroup(hs *service.Service) *HeyaHandleGroup {
	return &HeyaHandleGroup{hs: hs}
}
