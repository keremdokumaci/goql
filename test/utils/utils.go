package utils

import (
	"database/sql"
	"fmt"
	"testing"
)

type BaseSuite struct {
	DB *sql.DB
}

func (BaseSuite) SkipTestIfModeNot(t *testing.T, mode TestMode) {
	if GetTestMode() != mode {
		t.Skip("skipping tests for mode " + string(mode))
	}
}

func (b *BaseSuite) TearDownSuite() {
	b.DB.Close()
}

func (b *BaseSuite) PostgresConnection() error {
	host := "127.0.0.1"
	port := "54325"
	user := "postgres_test"
	password := "mysecretpassword"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	b.DB = db
	return nil
}
