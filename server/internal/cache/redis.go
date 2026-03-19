package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(hostname, port, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", hostname, port),
		Password:        "",
		DB:              0,
		PoolSize:        20,
		MinIdleConns:    5,
		MaxIdleConns:    10,
		ConnMaxIdleTime: 2 * time.Minute,
		ConnMaxLifetime: 10 * time.Minute,
	})
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := r.rdb.Get(ctx, key).Bytes()
	return value, err
}

func (r *RedisClient) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.rdb.Set(ctx, key, value, ttl).Err()
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}

func (r *RedisClient) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return r.rdb.Expire(ctx, key, ttl).Err()
}
