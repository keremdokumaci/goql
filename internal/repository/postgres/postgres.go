package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/keremdokumaci/goql/internal/models"
)

type postgresRepository[T models.Modeler] struct {
	db *sql.DB
}

func (r *postgresRepository[T]) Get(ctx context.Context, ID int) (T, error) {
	var model T
	tableName := model.TableName()
	query := fmt.Sprintf("select * from goql.%s where id=%d", tableName, ID)

	row := r.db.QueryRowContext(ctx, query)
	err := row.Scan(&model)
	if err != nil {
		return model, err
	}

	return model, nil
}

func New[T models.Modeler](db *sql.DB) *postgresRepository[T] {
	return &postgresRepository[T]{
		db: db,
	}
}
