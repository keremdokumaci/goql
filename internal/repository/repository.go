package repository

import (
	"context"
	"reflect"

	"github.com/keremdokumaci/goql/internal/models"
)

type Repository[T models.Modeler] interface {
	Get(ctx context.Context, ID int) (T, error)
	GetByUniqueField(ctx context.Context, field string, value any) (*T, error)
}

func GetFieldTypeByName[T models.Modeler](name string) reflect.Type {
	var bm T
	var reflectType reflect.Type
	t := reflect.TypeOf(bm)
	for i := 0; i < t.NumField(); i++ {
		dbField := t.Field(i).Tag.Get("db")
		if dbField == name {
			reflectType = t.Field(i).Type
			break
		}
	}

	return reflectType
}
