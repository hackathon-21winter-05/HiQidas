package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type FavoriteRepository interface {
	GetFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.History, error)
	CreateFavorite(ctx context.Context, history *model.History) error
	DeleteFavoriteByHeyaID(ctx context.Context, heyaID uuid.UUID) error
}
