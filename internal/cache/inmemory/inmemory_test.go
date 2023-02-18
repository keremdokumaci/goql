package inmemory

import (
	"testing"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Given
	key := "cache-key"
	val := "cache-value"

	// When
	ttlCache := ttlcache.New[string, any]()
	ttlCache.Set(key, val, DEFAULT_TTL)

	// Then
	sut := New(ttlCache)
	item := sut.Get(key)

	// Assertion
	assert.NotNil(t, item)
	assert.Equal(t, val, item.(string))
}

func TestSet(t *testing.T) {
	// Given
	key := "cache-key"
	val := "cache-value"

	// Then
	sut := New(nil)
	err := sut.Set(key, val, time.Second*10)

	// Assertion
	assert.Nil(t, err)
}
