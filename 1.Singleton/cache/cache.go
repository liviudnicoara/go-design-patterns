package cache

import "sync"

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
