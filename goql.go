package goql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	redisv9 "github.com/redis/go-redis/v9"

	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	"github.com/keremdokumaci/goql/internal/cache/redis"
	"github.com/keremdokumaci/goql/internal/cacher"
	"github.com/keremdokumaci/goql/internal/whitelist"
	whitelistrepository "github.com/keremdokumaci/goql/internal/whitelist/repository"
)

var (
	ErrUnexpectedDBType              error = errors.New("unexpected db type")
	ErrDBConfigurationIsMandatory    error = errors.New("db configuration is mandatory")
	ErrCacheConfigurationIsMandatory error = errors.New("cache configuration is mandatory")
)

type WhiteLister interface {
	QueryAllowed(ctx context.Context, queryName string) (bool, error)
}

type Cacher interface {
	CacheQuery(query string, response any, ttl ...time.Duration) error
	GetQueryCache(query string) any
}

type goQL struct {
	cache  cache.Cacher
	db     *sql.DB
	dbName constants.DB
}

// New returns a goQL struct pointer.
func New() *goQL {
	return &goQL{}
}

func (goql *goQL) ConfigureDB(dbName constants.DB, db *sql.DB) *goQL {
	goql.db = db
	goql.dbName = dbName

	return goql
}

func (goql *goQL) ConfigureInmemoryCache() *goQL {
	goql.cache = inmemory.New(nil)
	return goql
}

func (goql *goQL) ConfigureRedisCache(client *redisv9.Client) *goQL {
	goql.cache = redis.New(client)
	return goql
}

func (goql *goQL) UseWhitelister() (WhiteLister, error) {
	if goql.dbName == "" || goql.db == nil {
		return nil, ErrDBConfigurationIsMandatory
	}

	var repo whitelistrepository.WhitelistRepository
	repo, err := whitelistrepository.New(goql.dbName, goql.db)
	if err != nil {
		return nil, err
	}

	return whitelist.New(repo, goql.cache), nil
}

func (goql *goQL) UseGQLCacher() (Cacher, error) {
	if goql.cache == nil {
		return nil, ErrCacheConfigurationIsMandatory
	}

	return cacher.New(goql.cache), nil
}
