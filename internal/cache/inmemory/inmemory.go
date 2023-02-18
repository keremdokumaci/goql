package inmemory

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

const (
	DEFAULT_TTL = time.Minute * 10
)

type InmemoryCacher struct {
	cache *ttlcache.Cache[string, any]
}

func (c *InmemoryCacher) Get(key string) any {
	item := c.cache.Get(key)
	if item == nil {
		return nil
	}

	return item.Value()
}

func (c *InmemoryCacher) Set(key string, value any, ttl ...time.Duration) error {
	var timeToLive time.Duration = DEFAULT_TTL
	if len(ttl) > 0 {
		timeToLive = ttl[0]
	}

	c.cache.Set(key, value, timeToLive)
	return nil
}

func New(cache *ttlcache.Cache[string, any]) *InmemoryCacher {
	if cache == nil {
		cache = ttlcache.New[string, any]()
	}

	return &InmemoryCacher{
		cache: cache,
	}
}
