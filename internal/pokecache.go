package internal
import (
	"time"
	"sync"
)

type Cache struct {
	entry	map[string]cacheEntry
	m		sync.Mutex 
}

type cacheEntry struct {
	createdAt time.Time
	val		  []byte
}


func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	cache := &Cache{
		entry:  make(map[string]cacheEntry),
		m: 		sync.Mutex{},
	}
	go func () {
		for {
			<-ticker.C
			cache.reapLoop(interval)
		}
	}()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	if key == "" || len(val) == 0 {
		return
	}
	ce :=cacheEntry{
		createdAt: time.Now(),
		val: val}
	c.m.Lock()
	c.entry[key] = ce
	c.m.Unlock()
	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var rval []byte
	if key == "" {
		return rval, false
	}
	c.m.Lock()
	val, exists := c.entry[key]
	if !exists {
		c.m.Unlock()
		return rval, false
	}
	rval = val.val
	c.m.Unlock()
	return rval, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	now := time.Now()
	c.m.Lock()
	for key := range c.entry {
		if now.Sub(c.entry[key].createdAt) > interval {
			delete(c.entry, key)
		}
	}
	c.m.Unlock()
	return
}