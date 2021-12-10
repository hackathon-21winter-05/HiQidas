package utils

import (
	"context"
	"time"
)

func CreateTxContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 3*time.Second)
}
