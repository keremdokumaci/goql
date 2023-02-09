package postgres

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

func (p *postgresRepository) WhitelistExistsOperation(ctx context.Context, name string) bool {
	return p.db.QueryRowContext(ctx, "select * from goql.whitelists where name=$1", name) != nil
}

func New(db *sql.DB) *postgresRepository {
	return &postgresRepository{
		db: db,
	}
}
