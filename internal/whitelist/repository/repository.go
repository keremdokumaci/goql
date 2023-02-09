package repository

import (
	"context"
)

type Repository interface {
	WhitelistExistsOperation(ctx context.Context, name string) bool
}
