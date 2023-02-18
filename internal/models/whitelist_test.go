package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	whitelist := Whitelist{}
	tableName := whitelist.TableName()

	assert.Equal(t, "whitelists", tableName)
}
