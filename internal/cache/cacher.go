package cache

type cacher interface {
	Get(key string) any
	Set(key string, value any) error
}
