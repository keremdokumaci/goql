package redis

import (
	"context"
	"testing"
	"time"

	"github.com/keremdokumaci/goql/test/utils"
	"github.com/stretchr/testify/suite"
)

type RedisIntegrationTestSuite struct {
	utils.BaseSuite
	sut *RedisCacher
}

func (s *RedisIntegrationTestSuite) SetupSuite() {
	err := s.RedisConnection()
	if err != nil {
		s.T().Fatal(err.Error())
	}

	err = s.RedisCache.Set(context.Background(), "key", "val", time.Minute*5).Err()
	if err != nil {
		s.T().Fatal(err.Error())
	}

	s.sut = New(s.RedisCache)
}

func TestRedisIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(RedisIntegrationTestSuite))
}

func (s *RedisIntegrationTestSuite) TestGet() {
	// Given
	key := "key"

	// Then
	res := s.sut.Get(key)

	// Assertion
	s.NotNil(res)
	s.Equal("val", res)
}

func (s *RedisIntegrationTestSuite) TestGet_WithNoKey() {
	// Given
	key := "key_does_not_exist"

	// Then
	res := s.sut.Get(key)

	// Assertion
	s.Nil(res)
}

func (s *RedisIntegrationTestSuite) TestSet() {
	// Given
	key := "another_key"
	val := "val"

	// Then
	res := s.sut.Set(key, val, time.Second*5)

	// Assertion
	s.Nil(res)
	v := s.sut.Get(key)
	s.Equal(val, v)
}

func (s *RedisIntegrationTestSuite) TestSet_WithDefaultTTL() {
	// Given
	key := "another_key"
	val := "val"

	// Then
	res := s.sut.Set(key, val)

	// Assertion
	s.Nil(res)
	v := s.sut.Get(key)
	s.Equal(val, v)
}
