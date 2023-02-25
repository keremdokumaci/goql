package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/models"
	repositoryutils "github.com/keremdokumaci/goql/internal/repository/utils"
)

type postgresRepository[T models.Modeler] struct {
	sqlxDB    *sqlx.DB
	tableName string
}

func (r *postgresRepository[T]) Get(ctx context.Context, ID int) (T, error) {
	query := fmt.Sprintf("select * from goql.%s where id=%d", r.tableName, ID)

	var model T
	row := r.sqlxDB.QueryRowxContext(ctx, query)

	err := row.StructScan(&model)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (r *postgresRepository[T]) GetByUniqueField(ctx context.Context, field string, value any) (*T, error) {
	var model T

	query := fmt.Sprintf(
		`select * from "goql"."%s" where "%s"='%v'`,
		r.tableName,
		field,
		value,
	)

	row := r.sqlxDB.QueryRowxContext(ctx, query)
	err := row.StructScan(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func New[T models.Modeler](db *sql.DB) *postgresRepository[T] {
	var model T
	tableName := model.TableName()
	return &postgresRepository[T]{
		sqlxDB:    sqlx.NewDb(db, repositoryutils.GetDriverNameByDBName(constants.POSTGRES)),
		tableName: tableName,
	}
}
