package goql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	"github.com/keremdokumaci/goql/internal/cache/redis"
	"github.com/keremdokumaci/goql/internal/cacher"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository"
	"github.com/keremdokumaci/goql/internal/repository/postgres"
	"github.com/keremdokumaci/goql/internal/whitelist"
)

var (
	ErrUnexpectedDBType              error = errors.New("unexpected db type")
	ErrDBConfigurationIsMandatory    error = errors.New("db configuration is mandatory")
	ErrCacheConfigurationIsMandatory error = errors.New("cache configuration is mandatory")
)

type WhiteLister interface {
	OperationAllowed(ctx context.Context, operationName string) bool
}

type Cacher interface {
	GetOperation(operationName string) any
	CacheQuery(query string, response any, ttl ...time.Duration) error
}

type goQL struct {
	cache  cache.Cacher
	db     *sql.DB
	dbName DB
}

// New returns a goQL struct pointer.
func New() *goQL {
	return &goQL{}
}

func (goql *goQL) ConfigureDB(dbName DB, db *sql.DB) *goQL {
	goql.db = db
	goql.dbName = dbName

	return goql
}

func (goql *goQL) ConfigureCache(cacheName Cache) *goQL {
	switch cacheName {
	case INMEMORY:
		goql.cache = inmemory.New(nil)
	case REDIS:
		goql.cache = redis.New()
	}

	return goql
}

func (goql *goQL) UseWhitelister() (WhiteLister, error) {
	if goql.dbName == "" || goql.db == nil {
		return nil, ErrDBConfigurationIsMandatory
	}

	var repo repository.Repository[models.Whitelist]
	switch goql.dbName {
	case POSTGRES:
		repo = postgres.New[models.Whitelist](goql.db)
	default:
		return nil, ErrUnexpectedDBType
	}

	return whitelist.New(repo, goql.cache), nil
}

func (goql *goQL) UseGQLCacher() (Cacher, error) {
	if goql.cache == nil {
		return nil, ErrCacheConfigurationIsMandatory
	}

	return cacher.New(goql.cache), nil
}