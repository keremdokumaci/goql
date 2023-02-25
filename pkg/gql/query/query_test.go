package query

import (
	"testing"

	"github.com/keremdokumaci/goql/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	query, err := Parse(fixtures.Query)
	assert.Nil(t, err)
	assert.NotNil(t, query)
}

func TestParse_QueryHasNoName(t *testing.T) {
	query, err := Parse(fixtures.QueryWithNoName)
	assert.Nil(t, err)
	assert.NotNil(t, query)
}

func TestNewQuery_InvalidQuery(t *testing.T) {
	query, err := Parse(fixtures.InvalidQuery)
	assert.NotNil(t, err)
	assert.Nil(t, query)
}

func TestName(t *testing.T) {
	query, _ := Parse(fixtures.Query)
	queryName := query.Name()

	assert.Equal(t, "GetProductVariantByVariantIds", queryName)
}
