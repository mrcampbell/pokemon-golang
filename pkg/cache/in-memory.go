package cache

import (
	"fmt"

	"github.com/mrcampbell/pokemon-golang/pkg/app"
	gcache "github.com/patrickmn/go-cache"
)

type InMemoryCache struct {
	cache *gcache.Cache
}

func NewInMemoryCache() app.Cache {
	cache := gcache.New(gcache.NoExpiration, gcache.NoExpiration)
	return InMemoryCache{cache: cache}
}

func (c InMemoryCache) Get(key string) (string, error) {
	value, found := c.cache.Get(key)
	if !found {
		return "", fmt.Errorf("key not found")
	}
	return value.(string), nil
}

func (c InMemoryCache) Set(key string, value string) error {
	c.cache.Set(key, value, gcache.NoExpiration)
	return nil
}

func (c InMemoryCache) Delete(key string) error {
	c.cache.Delete(key)
	return nil
}
