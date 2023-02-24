package cacher

import (
	"testing"
	"time"

	cacheMock "github.com/keremdokumaci/goql/mocks/cache"
	"github.com/keremdokumaci/goql/test/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCacheQuery(t *testing.T) {
	// Given
	query := fixtures.Query
	response := "response"
	ttl := time.Second * 5

	// When
	cache := &cacheMock.Cacher{}
	cache.
		On(
			"Set",
			mock.AnythingOfType("string"),
			mock.Anything,
			mock.AnythingOfType("time.Duration")).
		Return(nil)

	// Then
	sut := New(cache)
	err := sut.CacheQuery(query, response, ttl)

	// Assertion
	assert.Nil(t, err)
}

func TestGetQueryCache(t *testing.T) {
	// Given
	query := fixtures.Query
	response := "response"

	// When
	cache := &cacheMock.Cacher{}
	cache.
		On(
			"Get",
			mock.AnythingOfType("string"),
		).
		Return(response)

	// Then
	sut := New(cache)
	res := sut.GetQueryCache(query)

	// Assertion
	assert.NotNil(t, res)
	assert.Equal(t, response, res)
}
