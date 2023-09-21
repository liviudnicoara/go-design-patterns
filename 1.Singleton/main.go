package main

import (
	"fmt"
	"sync"
)

// The Singleton pattern is a design pattern that restricts the instantiation of a class to a single instance.
// This is useful when exactly one object is needed to coordinate actions across the system.

var (
	once     sync.Once
	instance *CacheClient
)

type CacheClient struct {
	Server string
	Port   int
}

func GetCacheClient(server string, port int) *CacheClient {
	once.Do(func() {
		instance = &CacheClient{
			Server: server,
			Port:   port,
		}
	})

	return instance
}

func main() {
	cacheClient := GetCacheClient("192.168.1.1", 5000)

	fmt.Printf("Cache client settings: %+v\n", *cacheClient)

	cacheClient = GetCacheClient("192.168.2.2", 6000)

	fmt.Printf("Cache client settings: %+v\n", *cacheClient)
}
