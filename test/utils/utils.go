package utils

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type BaseSuite struct {
	DB         *sql.DB
	RedisCache *redis.Client
	suite.Suite
}

func (BaseSuite) SkipTestIfModeNot(t *testing.T, mode TestMode) { // nolint
	if GetTestMode() != mode {
		t.Skip("skipping tests for mode " + string(mode))
	}
}

func (b *BaseSuite) TearDownSuite() {
	if b.DB != nil {
		b.DB.Close()
	}
	if b.RedisCache != nil {
		b.RedisCache.Close()

	}
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

func (b *BaseSuite) RedisConnection() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := rdb.Ping(ctxTimeout).Err()
	if err != nil {
		return err
	}

	b.RedisCache = rdb
	return nil
}
