package heya

import "github.com/hackathon-21winter-05/HiQidas/repository"

type HeyaService interface {

}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func NewHeyaServiceImpl(repo repository.Repository) *HeyaServiceImpl {
	return &HeyaServiceImpl{repo: repo}
}


