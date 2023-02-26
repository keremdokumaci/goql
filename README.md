![Version](https://img.shields.io/badge/version-0.1.12-orange.svg)
[![GolangCI Lint](https://github.com/keremdokumaci/goql/actions/workflows/go-lint.yml/badge.svg)](https://github.com/keremdokumaci/goql/actions/workflows/go-lint.yml)
[![Golang Tests](https://github.com/keremdokumaci/goql/actions/workflows/go-test.yml/badge.svg)](https://github.com/keremdokumaci/goql/actions/workflows/go-test.yml)
![Test Coverage](https://img.shields.io/badge/coverage-91.9%25-orange.svg)

# GoQL

A GraphQL helper to deal with GraphQL requirements. Written with *Go 1.19*

## What does GoQL do ?

💾 `Caching`

🆘 `White Listing`

*And other features **will be** supported by GoQL.*

## Usage

You should create a **GoQL** instance at the beginning with using function

```go
gq := goql.New()
```

There are some functions for initialization of GoQL instance.

* ConfigureDB
* ConfigureInmemoryCache
* ConfigureRedisCache
* UseWhitelister
* UseGQLCacher

### Database Configuration

Database configuration is done by **ConfigureDB** function. It takes 2 parameters to configure database.

* Database name (Such as ***postgres***,***mysql***)
* Database instance (pointer of **sql.DB**)

🆘 For now, **postgres** is the only supported db. However, **mysql**, **mssql** and other most known database systems will be integrated to GoQL.

```go
gq.ConfigureDB(constants.POSTGRES, db)
```

#### Database Migration

Set your migration files directory as environment variable called **MIGRATION_DIR**. Then run

```go
migrations.MigratePostgres(db)
```

Required migration files can be found inside [migrations](./pkg/migrations/) folder. Migration files depends on database.

### Cache Configuration

Cache configuration is done by **ConfigureInmemoryCache** and **ConfigureRedisCache** functions.

```go
gq.ConfigureInmemoryCache()
```

```go
gq.ConfigureRedisCache(client)
```

### GQL Cacher

For caching GQL queries, you should run **UseGQLCacher** function. It returns **Cacher** and **error**.

🚨 *Database and cache configuration is **mandatory** to use this feature.*


```go
cacher,err := gq.UseGQLCacher()
```

### Whitelister

For whitelisting, you should run **UseWhitelister** function. It returns **Whitelister** interface and **error**.

🚨 *Database and cache configuration is **mandatory** to use this feature.*

🚨 *Database migration is mandatory to use whitelister*

### GQL Query Helper

```go
query, err := query.Parse("a gql query")
```

parses string query to [ast.Document]("github.com/graphql-go/graphql/language/ast").

There are some basic functions to get some specific data along query such as.

```go
queryName := query.Name()
```

There will be more different helper functions for GQL Operations.

## Testing

Run all tests

```sh
make test
```

Run single test

```sh
make test t="test_to_run"
```

Run code coverate

```sh
make test-coverage
```

Beside calculating code coverage, it replace code coverage badge on [README.md](./README.md).
