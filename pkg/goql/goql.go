package goql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/keremdokumaci/goql/internal/cache"
	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	"github.com/keremdokumaci/goql/internal/repository"
	"github.com/keremdokumaci/goql/internal/repository/postgres"
	"github.com/keremdokumaci/goql/internal/whitelist"
)

var (
	ErrUnexpectedDBType           error = errors.New("unexpected db type")
	ErrDBConfigurationIsMandatory error = errors.New("db configuration is mandatory")
)

/*
	TODO:
		- WhiteLister struct
			- there should be a migration mechanism to add & remove whitelisting dynamically.
		- Cacher struct
			- there should be a migration mechanism to add & remove whitelisting dynamically.
*/

type goQL struct {
	cache       cache.Cacher
	repository  repository.Repository
	whitelister whitelist.WhiteLister
}

// New returns a goQL struct pointer.
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

// ConfigureDB initializes repository via given dbName and db instance.
func (goql *goQL) ConfigureDB(dbName DB, db *sql.DB) error {
	// init repository
	switch dbName {
	case POSTGRES:
		goql.repository = postgres.New(db)
	default:
		return ErrUnexpectedDBType
	}

	// migrate database
	err := goql.repository.Migrate(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
