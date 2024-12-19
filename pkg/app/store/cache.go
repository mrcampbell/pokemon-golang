package store

type Cache interface {
	New() Cache
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
