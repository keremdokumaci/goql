package whitelist

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/keremdokumaci/goql/internal/cache/inmemory"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository/postgres"
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

	repository := postgres.New[models.Whitelist](s.DB)
	cache := inmemory.New(nil)

	s.sut = New(repository, cache)
}

func (s *WhitelistIntegrationTestSuite) BeforeTest(suiteName, testName string) {
	tx, _ := s.DB.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})

	_, err := tx.Exec(`INSERT INTO "goql"."whitelists" (operation_name, created_at) VALUES ($1,$2)`, "getProducts", time.Now())
	if err != nil {
		_ = tx.Rollback()
		s.T().Fatal(err.Error())
	}

	_ = tx.Commit()
}

func TestWhitelistIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(WhitelistIntegrationTestSuite))
}

func (s *WhitelistIntegrationTestSuite) TestOperationAllowed_OperationNotInCache() {
	// Given
	operationName := "getProducts"

	// When

	// Then
	allowed, err := s.sut.OperationAllowed(context.Background(), operationName)
	s.Nil(err)
	s.True(allowed)
}
