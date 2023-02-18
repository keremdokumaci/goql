package fixtures

var Query string = `
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

var QueryWithNoName string = `
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

var InvalidQuery string = `
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
