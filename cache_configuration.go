package goql

type Cache string

const (
	REDIS    Cache = "redis"
	INMEMORY Cache = "inmemory"
)
