package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type FavoriteRepository interface {
	GetFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Favorite, error)
	CreateFavorite(ctx context.Context, favorite *model.Favorite) error
	DeleteFavoriteByHeyaIDAndUserID(ctx context.Context, heyaID, userID uuid.UUID) error
	DeleteFavoriteByHeyaID(ctx context.Context, heyaID uuid.UUID) error
}
