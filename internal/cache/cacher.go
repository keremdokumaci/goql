package cache

import "time"

type Cacher interface {
	Get(key string) any
	Set(key string, value any, ttl ...time.Duration) error
}
