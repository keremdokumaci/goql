package whitelist

import (
	"context"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/whitelist/repository"
)

type WhiteLister interface {
	OperationAllowed(ctx context.Context, operationName string) bool
}

type whiteLister struct {
	repo   repository.Repository
	cacher cache.Cacher
}

func (w *whiteLister) OperationAllowed(ctx context.Context, operationName string) bool {
	cacheVal := w.cacher.Get(operationName)
	if cacheVal != nil {
		return true
	}

	exists := w.repo.WhitelistExistsOperation(ctx, operationName)
	if !exists {
		w.cacher.Set(operationName, true)
	}

	return exists
}

func New(repo repository.Repository, cacher cache.Cacher) WhiteLister {
	return &whiteLister{
		repo:   repo,
		cacher: cacher,
	}
}
