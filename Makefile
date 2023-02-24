.PHONY: all test clean
THIS_DIR:=$(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# Lint
lint:
	LOG_LEVEL=error golangci-lint run

# Migration
postgres_migration_dir=./internal/repository/migrations/postgres
mysql_migration_dir=./internal/repository/migrations/mysql
mssql_migration_dir=./internal/repository/migrations/mssql

create-migration:
	migrate create -ext sql -dir $postgres_migration_dir -seq $(name)
	migrate create -ext sql -dir $mysql_migration_dir -seq $(name)
	migrate create -ext sql -dir $mssql_migration_dir -seq $(name)

migrate:
	migrate -verbose -database $(database_url) -path migrations $(dir)

# Test
test:
	bash -c "$(THIS_DIR)scripts/test.sh $(ARGS)"
