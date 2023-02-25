package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	DEFAULT_TTL = time.Minute * 10
)

type RedisCacher struct {
	client *redis.Client
}

func (c *RedisCacher) Get(key string) any {
	val := c.client.Get(context.Background(), key).Val()
	if val == "" {
		return nil
	}

	return val
}

func (c *RedisCacher) Set(key string, value any, ttl ...time.Duration) error {
	var timeToLive time.Duration = DEFAULT_TTL
	if len(ttl) > 0 {
		timeToLive = ttl[0]
	}

	return c.client.Set(context.Background(), key, value, timeToLive).Err()
}

func New(client *redis.Client) *RedisCacher {
	return &RedisCacher{
		client: client,
	}
}
