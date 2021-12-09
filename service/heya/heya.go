package heya

import (
	"context"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
	"github.com/hackathon-21winter-05/HiQidas/service/utils"
)

type HeyaService interface {
	SaveHeya(title, description string) (*model.Heya, error)
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func (h *HeyaServiceImpl) SaveHeya(title, description string) (*model.Heya, error) {
	ctx, cancel := utils.CreateTxContext()
	defer cancel()

	err := h.repo.Do(ctx, nil, func(c context.Context) error {
			h.repo.
	})
}

func NewHeyaServiceImpl(repo repository.Repository) *HeyaServiceImpl {
	return &HeyaServiceImpl{repo: repo}
}
