package cache

import (
	"context"
	"time"
)

type Client interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, ttl time.Duration) error
}
