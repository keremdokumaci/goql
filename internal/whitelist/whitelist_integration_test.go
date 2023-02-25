package whitelist

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	whitelistrepository "github.com/keremdokumaci/goql/internal/whitelist/repository"
	"github.com/keremdokumaci/goql/pkg/migrations"
	testUtils "github.com/keremdokumaci/goql/test/utils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type WhitelistIntegrationTestSuite struct {
	sut *whiteLister
	testUtils.BaseSuite
}

func (s *WhitelistIntegrationTestSuite) SetupSuite() {
	os.Setenv("MIGRATION_DIR", "../../pkg/migrations/postgres")

	err := s.PostgresConnection()
	if err != nil {
		s.T().Fatal(err.Error())
	}

	err = migrations.MigratePostgres(s.DB)
	if err != nil {
		s.T().Fatal(err.Error())
	}

	repository, err := whitelistrepository.New(constants.POSTGRES, s.DB)
	if err != nil {
		s.T().Fatal(err.Error())
	}

	cache := inmemory.New(nil)

	s.sut = New(repository, cache)

	tx, _ := s.DB.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})

	queryID := 1

	_, err = tx.Exec(`INSERT INTO "goql"."queries" (id, name, created_at) VALUES ($1,$2,$3) RETURNING id`, queryID, "getProducts", time.Now())
	if err != nil {
		_ = tx.Rollback()
		s.T().Fatal(err.Error())
	}

	_, err = tx.Exec(`INSERT INTO "goql"."whitelists" (query_id, created_at) VALUES ($1,$2)`, queryID, time.Now())
	if err != nil {
		_ = tx.Rollback()
		s.T().Fatal(err.Error())
	}

	_ = tx.Commit()
}
func TestWhitelistIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(WhitelistIntegrationTestSuite))
}

func (s *WhitelistIntegrationTestSuite) TestQueryAllowed_QueryNotInCache() {
	// Given
	queryName := "getProducts"

	// Then
	allowed, err := s.sut.QueryAllowed(context.Background(), queryName)
	if err != nil {
		s.T().Fatal(err.Error())
	}

	// Assert
	s.Nil(err)
	s.True(allowed)
}

func (s *WhitelistIntegrationTestSuite) TestQueryAllowed_QueryNotInCacheAndDB() {
	// Given
	queryName := "xyz"

	// Then
	allowed, err := s.sut.QueryAllowed(context.Background(), queryName)
	if err != nil {
		s.T().Fatal(err.Error())
	}

	// Assert
	s.Nil(err)
	s.False(allowed)
}
