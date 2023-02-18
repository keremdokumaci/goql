package whitelist

import (
	"context"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository"
)

type WhiteLister interface {
	OperationAllowed(ctx context.Context, operationName string) bool
}

type whiteLister struct {
	repo   repository.Repository[models.Whitelist]
	cacher cache.Cacher
}

func (w *whiteLister) OperationAllowed(ctx context.Context, operationName string) bool {
	cacheVal := w.cacher.Get(operationName)
	if cacheVal != nil {
		return true
	}

	_, err := w.repo.Get(ctx, 123) //TODO: Get by specification.
	if err != nil {
		//TODO: log error.
		return false
	}

	return true
}

func New(repo repository.Repository[models.Whitelist], cacher cache.Cacher) WhiteLister {
	return &whiteLister{
		repo:   repo,
		cacher: cacher,
	}
}
