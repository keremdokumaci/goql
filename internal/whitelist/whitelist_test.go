package whitelist

import (
	"context"

	"github.com/keremdokumaci/goql/internal/models"
	cacheMock "github.com/keremdokumaci/goql/mocks/cache"
	repositoryMock "github.com/keremdokumaci/goql/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestWhitelistSuite struct {
	sut        WhiteLister
	cache      *cacheMock.Cacher
	repository *repositoryMock.Repository[models.Whitelist]
	suite.Suite
}

func (s *TestWhitelistSuite) BeforeTest() {
	s.cache = &cacheMock.Cacher{}
	s.repository = &repositoryMock.Repository[models.Whitelist]{}
}

func (s *TestWhitelistSuite) SetupSuite() {
	s.sut = New(s.repository, s.cache)
}

func (s *TestWhitelistSuite) TestOperationAllowed_OperationInCache() {
	// Given
	operationName := "operation_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(true)

	// Then
	allowed := s.sut.OperationAllowed(context.TODO(), operationName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get")
	s.repository.AssertNotCalled(s.T(), "Get")
	s.True(allowed)
}

func (s *TestWhitelistSuite) TestOperationAllowed_OperationNotInCache() {
	// Given
	operationName := "operation_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(nil)
	s.repository.On("Get", mock.AnythingOfType("context.Context"), mock.AnythingOfType("int"))

	// Then
	allowed := s.sut.OperationAllowed(context.TODO(), operationName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get")
	s.repository.AssertCalled(s.T(), "Get")
	s.True(allowed)
}
