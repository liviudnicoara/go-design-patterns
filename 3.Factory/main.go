package main

import (
	"fmt"

	"github.com/liviudnicoara/go-design-patterns/3.Factory/cache"
)

// The Factory is a pattern that provides an interface for creating objects in a superclass,
// but allows subclasses to alter the type of objects that will be created.
// This pattern is used when a class cannot anticipate the class of objects it needs to create.
// It is often used in situations where a system should be independent of how its objects are created,
// composed, and represented, providing a simple way of decoupling the concrete objects from the parts of the system
// that use them.

func main() {
	factory := cache.CacheFactory{}
	memCache, err := factory.CreateCache(cache.MemoryCacheType)

	if err != nil {
		fmt.Println(err)
	}

	disCache, err := factory.CreateCache(cache.DistributedCacheType)

	if err != nil {
		fmt.Println(err)
	}

	memCache.Set("m", "1")
	fmt.Println("Memory cache value: ", memCache.Get("m"))

	disCache.Set("d", "2")
	fmt.Println("Distributed cache value: ", disCache.Get("d"))
}
