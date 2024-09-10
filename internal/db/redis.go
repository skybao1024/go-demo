package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis 定义了 Redis 操作接口
type Redis interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
}

// RedisClient 是 Redis 客户端的实现
type RedisClient struct {
	client *redis.Client
}

// Set 实现 Redis 接口的 Set 方法
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

// Get 实现 Redis 接口的 Get 方法
func (r *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

// Del 实现 Redis 接口的 Del 方法
func (r *RedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Del(ctx, keys...)
}

// Exists 实现 Redis 接口的 Exists 方法
func (r *RedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Exists(ctx, keys...)
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}
