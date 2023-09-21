package main

import (
	"fmt"
)

// The Factory is a pattern that provides an interface for creating objects in a superclass,
// but allows subclasses to alter the type of objects that will be created.
// This pattern is used when a class cannot anticipate the class of objects it needs to create.
// It is often used in situations where a system should be independent of how its objects are created,
// composed, and represented, providing a simple way of decoupling the concrete objects from the parts of the system
// that use them.

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

func main() {
	factory := CacheFactory{}
	memCache, err := factory.CreateCache(MemoryCacheType)

	if err != nil {
		fmt.Println(err)
	}

	disCache, err := factory.CreateCache(DistributedCacheType)

	if err != nil {
		fmt.Println(err)
	}

	memCache.Set("m", "1")
	fmt.Println("Memory cache value: ", memCache.Get("m"))

	disCache.Set("d", "2")
	fmt.Println("Distributed cache value: ", disCache.Get("d"))
}
