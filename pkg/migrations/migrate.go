package migrations

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func GetLatestMigrationNumber(dir string) uint {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	lastEntry := entries[len(entries)-1].Name()
	version := strings.Split(lastEntry, "_")[0]

	versionNumber, err := strconv.Atoi(version)
	if err != nil {
		return 0
	}

	return uint(versionNumber)
}

func MigratePostgres(db *sql.DB, migration ...int) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+os.Getenv("MIGRATION_DIR"), "postgres", driver)
	if err != nil {
		return err
	}

	var migrationErr error
	if len(migration) > 0 {
		migrationErr = m.Steps(migration[0])
	} else {
		migrationErr = m.Up()
	}

	if migrationErr == migrate.ErrNoChange {
		return nil
	}

	return migrationErr
}
