package whitelist

import (
	"context"
	"database/sql"

	"github.com/keremdokumaci/goql/internal/cache"
	whitelistrepository "github.com/keremdokumaci/goql/internal/whitelist/repository"
)

type whiteLister struct {
	repo   whitelistrepository.WhitelistRepository
	cacher cache.Cacher
}

func (w *whiteLister) QueryAllowed(ctx context.Context, queryName string) (bool, error) {
	cacheVal := w.cacher.Get(queryName)
	if cacheVal != nil {
		return true, nil
	}

	_, err := w.repo.GetWhitelistByQueryName(queryName)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	_ = w.cacher.Set(queryName, true)
	return true, nil
}

func New(repo whitelistrepository.WhitelistRepository, cacher cache.Cacher) *whiteLister {
	return &whiteLister{
		repo:   repo,
		cacher: cacher,
	}
}
