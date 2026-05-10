package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop() {
	
}