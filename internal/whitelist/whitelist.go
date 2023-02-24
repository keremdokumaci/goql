package whitelist

import (
	"context"
	"database/sql"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository"
)

type whiteLister struct {
	repo   repository.Repository[models.Whitelist]
	cacher cache.Cacher
}

func (w *whiteLister) OperationAllowed(ctx context.Context, operationName string) (bool, error) {
	cacheVal := w.cacher.Get(operationName)
	if cacheVal != nil {
		return true, nil
	}

	_, err := w.repo.GetByUniqueField(ctx, "operation_name", operationName)
	if err == sql.ErrNoRows {
		return false, nil
	}

	return err == nil, err
}

func New(repo repository.Repository[models.Whitelist], cacher cache.Cacher) *whiteLister {
	return &whiteLister{
		repo:   repo,
		cacher: cacher,
	}
}
