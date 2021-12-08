package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type tsunaRepository interface {
	GetTsuna(ctx context.Context, id uuid.UUID) (*model.Tsuna, error)
	InsertTsuna(ctx context.Context, tsuna *model.Tsuna) error
	DeleteTsuna(ctx context.Context, tsuna *model.Tsuna) error
	UpdateTsuna(Ctx context.Context, tsuna *model.Tsuna) error
}
