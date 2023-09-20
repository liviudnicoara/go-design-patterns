package main

import (
	"fmt"

	"github.com/liviudnicoara/go-design-patterns/1.Singleton/cache"
)

// The Singleton pattern is a design pattern that restricts the instantiation of a class to a single instance.
// This is useful when exactly one object is needed to coordinate actions across the system.

func main() {
	cacheClient := cache.GetCacheClient("192.168.1.1", 5000)

	fmt.Printf("Cache client settings: %+v\n", *cacheClient)

	cacheClient = cache.GetCacheClient("192.168.2.2", 6000)

	fmt.Printf("Cache client settings: %+v\n", *cacheClient)
}
