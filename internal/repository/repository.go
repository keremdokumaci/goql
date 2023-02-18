package repository

import (
	"context"

	"github.com/keremdokumaci/goql/internal/models"
)

type Repository[T models.Modeler] interface {
	Get(ctx context.Context, ID int) (T, error)
}