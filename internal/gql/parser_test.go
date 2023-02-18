package gql

import (
	"testing"

	"github.com/keremdokumaci/goql/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	doc, err := Parse(fixtures.Query)

	assert.Equal(t, doc.OperationName(), "GetProductVariantByVariantIds")
	assert.Nil(t, err)
}

func TestParse_QueryHasNoName(t *testing.T) {
	doc, err := Parse(fixtures.QueryWithNoName)

	assert.Equal(t, doc.OperationName(), "")
	assert.Nil(t, err)
}
func TestParse_InvalidQuery(t *testing.T) {
	doc, err := Parse(fixtures.InvalidQuery)

	assert.Nil(t, doc)
	assert.NotNil(t, err)
}
