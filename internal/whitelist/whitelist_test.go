package whitelist

import (
	"context"
	"testing"
	"time"

	"github.com/keremdokumaci/goql/internal/models"
	cacheMock "github.com/keremdokumaci/goql/mocks/cache"
	repositoryMock "github.com/keremdokumaci/goql/mocks/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WhitelistTestSuite struct {
	sut        WhiteLister
	cache      *cacheMock.Cacher
	repository *repositoryMock.Repository[models.Whitelist]
	suite.Suite
}

func (s *WhitelistTestSuite) BeforeTest(suiteName, testName string) {
	s.cache = &cacheMock.Cacher{}
	s.repository = &repositoryMock.Repository[models.Whitelist]{}
	s.sut = New(s.repository, s.cache)
}

func TestWhitelistTestSuite(t *testing.T) {
	suite.Run(t, new(WhitelistTestSuite))
}

func (s *WhitelistTestSuite) TestOperationAllowed_OperationInCache() {
	// Given
	operationName := "operation_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(true)

	// Then
	allowed := s.sut.OperationAllowed(context.TODO(), operationName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get", mock.AnythingOfType("string"))
	s.repository.AssertNotCalled(s.T(), "Get")
	s.True(allowed)
}

func (s *WhitelistTestSuite) TestOperationAllowed_OperationNotInCache() {
	// Given
	operationName := "operation_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(nil)
	s.repository.On("GetByUniqueField", mock.Anything, mock.AnythingOfType("string"), mock.Anything).
		Return(&models.Whitelist{
			OperationName: "test_operation",
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now(),
			},
		}, nil)

	// Then
	allowed := s.sut.OperationAllowed(context.TODO(), operationName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get", operationName)
	s.repository.AssertCalled(s.T(), "GetByUniqueField", context.TODO(), "operation_name", operationName)
	s.True(allowed)
}
