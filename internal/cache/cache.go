package cache

import (
	"context"
	"time"

	"github.com/phongnd2802/daily-social/internal/global"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
	SetEx(ctx context.Context, key string, value interface{}, expirationTimeSecond int64) error
	TTL(ctx context.Context, key string) (time.Duration, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (int64, error)
}

type redisCache struct {
	client *redis.Client
}

// Exists implements Cache.
func (r *redisCache) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}

// Del implements Cache.
func (r *redisCache) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// TTL implements Cache.
func (r *redisCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}

// SetEx implements Cache.
func (r *redisCache) SetEx(ctx context.Context, key string, value interface{}, expirationTimeSecond int64) error {
	return r.client.Set(ctx, key, value, time.Duration(expirationTimeSecond)*time.Second).Err()
}

// Get implements Cache.
func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Set implements Cache.
func (r *redisCache) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func NewRedisClient() Cache {
	return &redisCache{
		client: global.Rdb,
	}
}
