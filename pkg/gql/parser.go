package gql

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/pkg/errors"
)

type doc struct {
	*ast.Document
}

func (d *doc) OperationName() string {
	definition := d.Definitions[0]

	operationDef, ok := definition.(*ast.OperationDefinition)
	if !ok || operationDef.Name == nil {
		return ""
	}

	return operationDef.Name.Value
}

func Parse(query string) (*doc, error) {
	document, err := parser.Parse(parser.ParseParams{
		Source: query,
	})

	if err != nil {
		return nil, errors.Wrap(err, "[gql].[parse]")
	}

	return &doc{
		Document: document,
	}, nil
}
