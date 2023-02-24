package cacher

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/keremdokumaci/goql/internal/cache"
)

type gqlCacher struct {
	cache cache.Cacher
}

func New(cache cache.Cacher) *gqlCacher {
	return &gqlCacher{
		cache: cache,
	}
}

func hashQuery(query string) string {
	hashAlgorithm := sha256.New()
	hashAlgorithm.Write([]byte(query))
	return hex.EncodeToString(hashAlgorithm.Sum(nil))
}

func (g *gqlCacher) CacheQuery(query string, response any, ttl ...time.Duration) error {
	var tl time.Duration
	if len(ttl) > 0 {
		tl = ttl[0]
	}

	return g.cache.Set(hashQuery(query), response, tl)
}

func (g *gqlCacher) GetQueryCache(query string) any {
	return g.cache.Get(hashQuery(query))
}
