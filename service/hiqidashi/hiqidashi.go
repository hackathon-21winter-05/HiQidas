package hiqidashi

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type HiqidashiService interface {
	GetHiqidashisByHeyaID(c context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error)
	CreateHiqidashi(c context.Context, createrID, heyaID, parentID uuid.UUID) (*model.Hiqidashi, error)
	UpdateHiqidashiByID(c context.Context, hiqidashi *model.NullHiqidashi) error
}

type HiqidashiServiceImpl struct {
	repo repository.Repository
}

func NewHiqidashiService(repo repository.Repository) HiqidashiService {
	return &HiqidashiServiceImpl{
		repo: repo,
	}
}
