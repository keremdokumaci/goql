package main

import (
	"github.com/keremdokumaci/goql/pkg/goql"
)

func main() {
	gq := goql.New()
	gq.ConfigureCache(goql.INMEMORY)
	cacher, _ := gq.UseGQLCacher()
	cacher.GetOperation("getProducts")
}
