package heya

import (
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type HeyaService interface {
	SaveHeya(title,description string) (*model.Heya,error)
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func NewHeyaServiceImpl(repo repository.Repository) *HeyaServiceImpl {
	return &HeyaServiceImpl{repo: repo}
}


