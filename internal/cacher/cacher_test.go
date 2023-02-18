package cacher

import (
	"testing"
	"time"

	cacheMock "github.com/keremdokumaci/goql/mocks/cache"
	"github.com/keremdokumaci/goql/test/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOperation(t *testing.T) {
	// Given
	operationName := "test-operation"

	// When
	cache := &cacheMock.Cacher{}
	cache.On("Get", mock.AnythingOfType("string")).Return("query-response")

	// Then
	sut := New(cache)
	response := sut.GetOperation(operationName)

	// Assertion
	assert.NotNil(t, response)
}

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
