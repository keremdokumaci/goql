package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/stretchr/testify/suite"
)

type testModel struct {
	models.BaseModel
}

func (testModel) TableName() string {
	return "test_models"
}

type PostgresRepositorySuite struct {
	suite.Suite
	sqlMock sqlmock.Sqlmock
	db      *sql.DB
	sut     *postgresRepository[testModel]
}

func (s *PostgresRepositorySuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s.sqlMock = mock
	s.db = db

	s.sut = New[testModel](s.db)
}

func (s *PostgresRepositorySuite) TearDownSuite() {
	s.db.Close()
}

func TestPostgresRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresRepositorySuite))
}

func (s *PostgresRepositorySuite) TestGet() {
	// Given
	ctx := context.TODO()
	id := 123

	// When
	expectedQuery := fmt.Sprintf("select * from goql.test_models where id=%d", id)
	s.sqlMock.
		ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(
			sqlmock.
				NewRows([]string{"id", "created_at", "updated_at"}).
				AddRow(id, time.Now(), nil),
		)

	// Then
	model, err := s.sut.Get(ctx, id)

	// Assertion
	s.Nil(err)
	s.Equal(id, model.ID)
}

func (s *PostgresRepositorySuite) TestGetByUniqueField() {
	// Given
	ctx := context.TODO()
	value := "get"
	field := "operation_name"

	// When
	expectedQuery := fmt.Sprintf(`select * from "goql"."test_models" where "%s"='%v'`, field, value)
	s.sqlMock.
		ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(
			sqlmock.
				NewRows([]string{"id", "created_at", "updated_at"}).
				AddRow(1, time.Now(), nil),
		)

	// Then
	model, err := s.sut.GetByUniqueField(ctx, field, value)

	// Assertion
	s.Nil(err)
	s.NotNil(model)
}
