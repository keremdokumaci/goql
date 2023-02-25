package query

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

type Query interface {
	Name() string
}

type query struct {
	*ast.Document
}

func (q *query) Name() string {
	definition := q.Definitions[0]

	operationDef, ok := definition.(*ast.OperationDefinition)
	if !ok || operationDef.Name == nil {
		return ""
	}

	return operationDef.Name.Value
}

func Parse(q string) (Query, error) {
	document, err := parser.Parse(parser.ParseParams{
		Source: q,
	})

	if err != nil {
		return nil, err
	}

	return &query{Document: document}, nil
}
