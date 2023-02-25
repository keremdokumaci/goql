package whitelist

import (
	"context"
	"testing"
	"time"

	"github.com/keremdokumaci/goql/internal/models"
	cacheMock "github.com/keremdokumaci/goql/mocks/cache"
	whitelistRepositoryMock "github.com/keremdokumaci/goql/mocks/whitelist/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WhitelistTestSuite struct {
	sut                 *whiteLister
	cache               *cacheMock.Cacher
	whitelistRepository *whitelistRepositoryMock.WhitelistRepository
	suite.Suite
}

func (s *WhitelistTestSuite) BeforeTest(suiteName, testName string) {
	s.cache = &cacheMock.Cacher{}
	s.whitelistRepository = &whitelistRepositoryMock.WhitelistRepository{}
	s.sut = New(s.whitelistRepository, s.cache)
}

func TestWhitelistTestSuite(t *testing.T) {
	suite.Run(t, new(WhitelistTestSuite))
}

func (s *WhitelistTestSuite) TestQueryAllowed_QueryInCache() {
	// Given
	queryName := "query_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(true)

	// Then
	allowed, err := s.sut.QueryAllowed(context.TODO(), queryName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get", mock.AnythingOfType("string"))
	s.whitelistRepository.AssertNotCalled(s.T(), "GetWhitelistByQueryName")
	s.True(allowed)
	s.Nil(err)
}

func (s *WhitelistTestSuite) TestQueryAllowed_QueryNotInCache() {
	// Given
	queryName := "query_name"

	// When
	s.cache.On("Get", mock.AnythingOfType("string")).Return(nil)
	s.cache.On("Set", mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(nil)
	s.whitelistRepository.On("GetWhitelistByQueryName", mock.AnythingOfType("string")).
		Return(&models.Whitelist{
			QueryID: 1,
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now(),
			},
		}, nil)

	// Then
	allowed, err := s.sut.QueryAllowed(context.TODO(), queryName)

	// Assertion
	s.cache.AssertCalled(s.T(), "Get", queryName)
	s.whitelistRepository.AssertCalled(s.T(), "GetWhitelistByQueryName", queryName)
	s.True(allowed)
	s.Nil(err)
}
