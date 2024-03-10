package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createAt time.Time
	value    []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
	}

	go newCache.ExpireCache(interval)

	return newCache
}

func (c *Cache) Add(key string, data []byte) {
	c.cache[key] = cacheEntry{
		createAt: time.Now(),
		value:    data,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	v, ok := c.cache[key]

	return v.value, ok
}

func (c *Cache) ExpireCache(interval time.Duration) {
	intervalTime := time.Now().Add(-interval)
	ticker := time.NewTicker(interval)

	for range ticker.C {
		for k, v := range c.cache {
			if v.createAt.Before(intervalTime) {
				delete(c.cache, k)
			}
		}
	}
}
