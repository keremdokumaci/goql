package repositoryutils

import (
	"reflect"
	"testing"

	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetDriverNameByDBName_Posgres(t *testing.T) {
	driver := GetDriverNameByDBName(constants.POSTGRES)
	assert.Equal(t, "pgx", driver)
}

func TestGetDriverNameByDBName_Unknown(t *testing.T) {
	driver := GetDriverNameByDBName("xyz")
	assert.Equal(t, "", driver)
}

func TestGetFieldTypeByName(t *testing.T) {
	fieldType := GetFieldTypeByName[models.Whitelist]("query_id")
	assert.IsType(t, reflect.Int, fieldType.Kind())
}
