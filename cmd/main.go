package main

import (
	"fmt"

	"github.com/keremdokumaci/goql/pkg/gql"
)

func main() {
	doc, err := gql.Parse("query")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doc.OperationName())
}
