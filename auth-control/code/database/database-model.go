package database

import (
	appErrors "auth-control/errors"
	"context"
	"time"
)

type IDatabase interface {
	Set(ctx context.Context, key string, value string, ex time.Duration) (string, appErrors.ErrorResponse)
	Get(ctx context.Context, key string) (string, appErrors.ErrorResponse)
	Delete(ctx context.Context, key string) appErrors.ErrorResponse
}

var Database IDatabase
