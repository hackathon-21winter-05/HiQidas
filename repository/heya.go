package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HeyaRepository interface {
	GetHeyasID(ctx context.Context) ([]uuid.UUID, error)
	GetHeyaByID(ctx context.Context, id uuid.UUID) (*model.Heya, error)
	CreateHeya(ctx context.Context, title, description string) (*model.Heya, error)
	UpdateHeya(ctx context.Context, heya *model.Heya) error
	DeleteHeya(ctx context.Context, id uuid.UUID) error
}
