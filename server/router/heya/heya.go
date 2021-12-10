package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/service/heya"
)

type HeyaHandleGroup struct {
	hs heya.HeyaService
}

func NewHeyaHandleGroup(hs heya.HeyaService) *HeyaHandleGroup {
	return &HeyaHandleGroup{hs: hs}
}
