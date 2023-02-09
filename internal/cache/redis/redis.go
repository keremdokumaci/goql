package redis

import "time"

type RedisCacher struct {
}

func (c *RedisCacher) Get(key string) any {
	return nil
}

func (c *RedisCacher) Set(key string, value any, ttl ...time.Duration) error {
	return nil
}

func (c *RedisCacher) Options() error {
	return nil
}

func New() *RedisCacher {
	return &RedisCacher{}
}
