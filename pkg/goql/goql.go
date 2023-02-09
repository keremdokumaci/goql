package goql

import (
	"database/sql"
	"errors"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	"github.com/keremdokumaci/goql/internal/whitelist"
	"github.com/keremdokumaci/goql/internal/whitelist/repository"
	"github.com/keremdokumaci/goql/internal/whitelist/repository/postgres"
)

var (
	ErrUnexpectedDBType           error = errors.New("unexpected db type")
	ErrDBConfigurationIsMandatory error = errors.New("db configuration is mandatory")
)

type goQL struct {
	CacheAll bool

	whiteList []string
	cacheList []string

	cache       cache.Cacher
	repository  repository.Repository
	whitelister whitelist.WhiteLister
}

func New() *goQL {
	var goql *goQL
	goql.cache = inmemory.New()
	return goql
}

func (goql *goQL) UseWhitelist(operationNames ...string) error {
	if goql.repository == nil {
		return ErrDBConfigurationIsMandatory
	}
	goql.whitelister = whitelist.New(goql.repository, goql.cache)
	return nil
}

func (goql *goQL) ConfigureDB(dbName DB, db *sql.DB) error {
	switch dbName {
	case POSTGRES:
		goql.repository = postgres.New(db)
	default:
		return ErrUnexpectedDBType
	}

	return nil
}
