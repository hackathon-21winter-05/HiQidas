package repository

import (
	"context"
	"github.com/hackathon-21winter-05/HiQidas/model"
)

type HistoryRepository interface {
	GetHistoriesUserID(ctx context.Context)(*[]model.User,error)
	GetHistoriesHeyaID()
}
