package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository/postgres"
)

var (
	ErrUnexpectedDBType error = errors.New("unexpected db type")
)

type Repository[T models.Modeler] interface {
	Get(ctx context.Context, ID int) (T, error)
	GetByUniqueField(ctx context.Context, field string, value any) (*T, error)
}

func NewRepository[T models.Modeler](dbName constants.DB, db *sql.DB) (Repository[T], error) {
	switch dbName {
	case constants.POSTGRES:
		return postgres.New[T](db), nil
	default:
		return nil, ErrUnexpectedDBType
	}
}
