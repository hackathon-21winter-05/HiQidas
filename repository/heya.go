package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HeyaRepository interface {
	GetHeyas(ctx context.Context) ([]*model.Heya, error)
	GetHeyaByID(ctx context.Context, id uuid.UUID) (*model.Heya, error)
	GetHeyasByCreatorID(ctx context.Context, creatorID uuid.UUID) ([]*model.Heya,error)
	CreateHeya(ctx context.Context, heya *model.Heya) error
	UpdateHeyaByID(ctx context.Context, heya *model.NullHeya) error
	DeleteHeyaByID(ctx context.Context, id uuid.UUID) error
}
