package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HistoryRepository interface {
	GetHistoriesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.History, error)
	CreateHistory(ctx context.Context, userID uuid.UUID) error
	DeleteHistory(ctx context.Context, history *model.History) error
}
