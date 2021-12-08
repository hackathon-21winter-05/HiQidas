package repository

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HistoryRepository interface {
	GetHistories(ctx context.Context) ([]*model.History, error)
	GetHistoriesByUserID(ctx context.Context, userID uuid.UUID) ([]*model.History, error)
	CreateHistory(ctx context.Context, heyaID uuid.UUID) error //部屋が作成されたらこれも叩く
	DeleteHistory(ctx context.Context, heyaID uuid.UUID) error
}
