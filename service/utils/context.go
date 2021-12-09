package utils

import (
	"context"
	"time"
)

func CreateTxContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
