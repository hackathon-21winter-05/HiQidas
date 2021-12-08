package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type TsunaRepository interface {
	GetTsunasByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Tsuna, error)
	CreateTsuna(ctx context.Context, tsuna *model.Tsuna) error
	DeleteTsuna(ctx context.Context, id uuid.UUID) error
	UpdateTsuna(Ctx context.Context, tsuna *model.Tsuna) error
}
