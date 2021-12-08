package repository

import (
	"context"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HistoryRepository interface {
	GetHistoriesByUserID(ctx context.Context) ([]*model.History, error)
	DeleteHistory(ctx context.Context, heyaId context.Context) error
}
