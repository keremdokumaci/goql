package migrations

import (
	"testing"

	"github.com/keremdokumaci/goql/test/utils"
	"github.com/stretchr/testify/suite"
)

type MigrateTestSuite struct {
	suite.Suite
	utils.BaseSuite
}

func (s *MigrateTestSuite) BeforeTest(suiteName, testName string) {
	// s.SkipTestIfModeNot(s.T(), utils.INTEGRATION)
}

func (s *MigrateTestSuite) SetupSuite() {
	err := s.PostgresConnection()
	if err != nil {
		s.T().Fatal(err.Error())
	}
}

func TestMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateTestSuite))
}

func (s *MigrateTestSuite) TestMigratePostgres() {
	err := MigratePostgres(s.DB)
	s.Nil(err)

	latestVersion := GetLatestMigrationNumber("./postgres")

	row := s.DB.QueryRow("select version from public.schema_migrations")
	s.Nil(row.Err())

	var version uint
	row.Scan(&version)

	s.Equal(latestVersion, version)
}
