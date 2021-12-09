package utils

import (
	"context"
	"time"
)

func createTxContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
