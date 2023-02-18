package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keremdokumaci/goql/internal/models"
)

type postgresRepository[T models.Modeler] struct {
	sqlxDB *sqlx.DB
}

func (r *postgresRepository[T]) Get(ctx context.Context, ID int) (T, error) {
	var model T
	tableName := model.TableName()
	query := fmt.Sprintf("select * from goql.%s where id=%d", tableName, ID)

	row := r.sqlxDB.QueryRowxContext(ctx, query)
	err := row.StructScan(&model)
	if err != nil {
		return model, err
	}

	return model, nil
}

func New[T models.Modeler](db *sql.DB) *postgresRepository[T] {
	return &postgresRepository[T]{
		sqlxDB: sqlx.NewDb(db, "pgx"),
	}
}
