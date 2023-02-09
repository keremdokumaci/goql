package repository

import (
	"context"
)

type Repository interface {
	Migrate(ctx context.Context) error
	WhitelistExistsOperation(ctx context.Context, name string) bool
}
