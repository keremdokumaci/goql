package gql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	query := `
		query GetProductVariantByVariantIds($first: Int!, $ids:[ID!]) {
			productVariants(ids: $ids, first: $first) {
				edges {
				node {
					id
					name
					sku
					stocks {
					warehouse {
						id
					}
					}
					product {
						id
						name
					}
				}
				}
		}
		}
	`

	doc, err := Parse(query)

	assert.Equal(t, doc.OperationName(), "GetProductVariantByVariantIds")
	assert.Nil(t, err)
}

func TestParse_QueryHasNoName(t *testing.T) {
	query := `
		{
			productVariants(ids: $ids, first: $first) {
				edges {
				node {
					id
					name
					sku
					stocks {
					warehouse {
						id
					}
					}
					product {
						id
						name
					}
				}
				}
		}
		}
	`

	doc, err := Parse(query)

	assert.Equal(t, doc.OperationName(), "")
	assert.Nil(t, err)
}
func TestParse_InvalidQuery(t *testing.T) {
	query := `
		{
			(ids: $ids, first: $first) {
				edges {
				node {
					id
					name
					sku
					stocks {
					warehouse {
						id
					}
					}
					product {
						id
						name
					}
				}
				}
		}
		}
	`

	doc, err := Parse(query)

	assert.Nil(t, doc)
	assert.NotNil(t, err)
}
