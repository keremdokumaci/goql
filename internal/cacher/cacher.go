package cacher

import "github.com/keremdokumaci/goql/internal/cache"

type GQLCacher interface {
	GetOperation(operationName string) (string, error)
}

type gqlCacher struct {
	cache cache.Cacher
}

func New(cache cache.Cacher) GQLCacher {
	return &gqlCacher{
		cache: cache,
	}
}

func (g *gqlCacher) GetOperation(operationName string) (string, error) {
	return "", nil
}
