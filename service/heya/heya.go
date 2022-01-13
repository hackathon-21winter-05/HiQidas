package heya

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
	"github.com/hackathon-21winter-05/HiQidas/repository"
)

type HeyaService interface {
	CreateHeya(c context.Context, userID uuid.UUID, title, description string) (*model.Heya, error)
	DeleteHeya(c context.Context, heyaID uuid.UUID) error
	GetHeyas(c context.Context) ([]*model.Heya, error)
	GetHeyaByID(c context.Context, heyaID uuid.UUID) (*model.Heya, error)
	PutHeyaByID(c context.Context, heya *model.NullHeya, heyaID, userID uuid.UUID) error
	PutFavoriteByHeyaID(c context.Context, heyaID uuid.UUID, userID uuid.UUID, isFavorite bool) error
}

type HeyaServiceImpl struct {
	repo repository.Repository
}

func NewHeyaService(repo repository.Repository) HeyaService {
	return &HeyaServiceImpl{
		repo: repo,
	}
}
