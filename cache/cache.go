/*

Package cache implements a naïve in-memory cache, safe for concurrent
access.  Largely adapted from https://golang.org/x/build/internal/lru.

Examples

This creates a new cache that can hold 50 items:

	c := cache.New(50)

Add something to this cache; if the cache contains more than 50 items,
the oldest item is removed to make space:

	c.Add("foo", "bar")

Get something from this cache:

	foo := c.Get("foo")

As a naïve cache, there are no methods for clearing or deleting items,
even for checking for the current number of items in the cache; for more
sophisticated caching, see https://github.com/golang/groupcache or
https://github.com/hashicorp/golang-lru.

*/
package cache

import (
	"container/list"
	"sync"
)

// Cache is a naive in-memory cache.  Safe for concurrent access.
type Cache struct {
	// MaxEntries is the maximum number of entries before an item is
	// evicted.  Zero means no limit.
	MaxEntries int

	mu    sync.Mutex
	ll    *list.List
	cache map[interface{}]*list.Element
}

// Key may be any value that is comparable.
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit.
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key Key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if _, ok := c.cache[key]; ok {
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		old := c.ll.Back()
		if old != nil {
			c.ll.Remove(old)
			kv := old.Value.(*entry)
			delete(c.cache, kv.key)
		}
	}
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		return ele.Value.(*entry).value, true
	}
	return
}
