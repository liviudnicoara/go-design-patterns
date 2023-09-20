package cache

import (
	"fmt"
)

type Cache interface {
	Set(string, string)
	Get(string) string
}

type MemoryCache struct {
	storage map[string]string
}

func (c *MemoryCache) Set(key, value string) {
	c.storage[key] = value
}

func (c *MemoryCache) Get(key string) string {
	return c.storage[key]
}

type DistributedCache struct {
	storage map[string]string
}

func (c *DistributedCache) Set(key, value string) {
	c.storage[key] = value
}

func (c *DistributedCache) Get(key string) string {
	return c.storage[key]
}

type CacheType string

const (
	MemoryCacheType      CacheType = "memory"
	DistributedCacheType CacheType = "distributed"
)

type CacheFactory struct{}

func (lf *CacheFactory) CreateCache(cacheType CacheType) (Cache, error) {
	switch cacheType {
	case MemoryCacheType:
		return &MemoryCache{storage: map[string]string{}}, nil
	case DistributedCacheType:
		return &DistributedCache{storage: map[string]string{}}, nil
	default:
		return nil, fmt.Errorf("unsupported cache type")
	}
}
