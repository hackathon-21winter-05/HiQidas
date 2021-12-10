package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HiqidashiRepository interface {
	GetHiqidashisByID(ctx context.Context, id uuid.UUID) (*model.Hiqidashi, error)
	GetHiqidashisByHeyaID(ctx context.Context, heyaID uuid.UUID) ([]*model.Hiqidashi, error)
	GetHiqidashisByParentID(ctx context.Context, parentID uuid.UUID) ([]*model.Hiqidashi, error)
	CreateHiqidashi(ctx context.Context, hiqidashi *model.Hiqidashi) error
	DeleteHiqidashiByID(ctx context.Context, id uuid.UUID) error
	UpdateHiqidashiByID(ctx context.Context, hiqidashi *model.NullHiqidashi) error
	DeleteHiqidashiDrawing(ctx context.Context, hiqidashi *model.Hiqidashi) error
}
