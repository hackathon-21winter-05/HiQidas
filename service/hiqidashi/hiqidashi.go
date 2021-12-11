package hiqidashi

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

var defaultColors = []string{"0F2540", "00AA90", "E03C8A", "9E7A7A", "2EA9DF", "42602D", "77428D", "FFB11B"}

type HiqidashiService interface {
	GetHiqidashis(c context.Context) ([]*model.Hiqidashi, error)
	GetHiqidashisByHeyaID(c context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error)
	CreateHiqidashi(c context.Context, createrID, heyaID, parentID uuid.UUID) (*model.Hiqidashi, error)
	UpdateHiqidashiByID(c context.Context, hiqidashi *model.NullHiqidashi) error
	DeleteHiqidashiByID(c context.Context, id uuid.UUID) error
}

type HiqidashiServiceImpl struct {
	repo repository.Repository
}

func NewHiqidashiService(repo repository.Repository) HiqidashiService {
	return &HiqidashiServiceImpl{
		repo: repo,
	}
}
