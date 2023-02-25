package repositoryutils

import (
	"reflect"

	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/models"
)

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

func GetDriverNameByDBName(dbName constants.DB) string {
	switch dbName {
	case constants.POSTGRES:
		return "pgx"
	default:
		return ""
	}
}
