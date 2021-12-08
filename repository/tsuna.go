package repository

import (
	"context"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type tsunaRepository interface {
	GetTsuna(ctx context.Context) (*model.Tsuna, error)
	InsertTsuna(ctx context.Context, tsuna *model.Tsuna) error
	DeleteTsuna(ctx context.Context,tsuna *model.Tsuna) error
}
