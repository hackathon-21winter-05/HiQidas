package repository

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HeyaRepository interface {
	GetHeyasID(ctx context.Context) ([]uuid.UUID, error)
	GetHeyaByID(ctx context.Context, id uuid.UUID) (*model.Heya, error)
	GetHeyaByIDUsers(ctx context.Context, heyaID uuid.UUID) (*model.User, error)
	InsertHeya(ctx context.Context, title string, description sql.NullString) (*model.Heya, error)
	UpdateHeyaTitle(ctx context.Context, title string) error
	UpdateHeyaDescription(ctx context.Context, description string) error
	DeleteHeya(ctx context.Context, id uuid.UUID) error
}
