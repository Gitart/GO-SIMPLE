# Quick tip: Implementing an in-memory cache in Go

**Published on:** December 22nd, 2023

In almost all web applications that I build, I end up needing to persist some data – either for a short period of time (such as caching the result of an expensive database query), or for the lifetime of the running application until it is restarted.

When your application is a single binary running on a single machine, a simple, effective, and no-dependency way to do this is by persisting the data in memory using a mutex-protected map. And since generics was introduced in Go 1.18, it's possible to write a generic implementation that you can use to persist various different data types in a type-safe way.

**Note:** The code for this post can be found in [this gist](https://gist.github.com/alexedwards/af9bc1a31964342c199e9a832ab91b77).

## Long-lived cache

If you want to persist data for the lifetime of the running application (or until you deliberately delete the data), you can create a generic `Cache` type like this:

```go
package cache

import (
    "sync"
    "time"
)

// Cache is a basic in-memory key-value cache implementation.
type Cache[K comparable, V any] struct {
    items map[K]V     // The map storing key-value pairs.
    mu    sync.Mutex  // Mutex for controlling concurrent access to the cache.
}

// New creates a new Cache instance.
func New[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        items: make(map[K]V),
    }
}

// Set adds or updates a key-value pair in the cache.
func (c *Cache[K, V]) Set(key K, value V) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = value
}

// Get retrieves the value associated with the given key from the cache. The bool
// return value will be false if no matching key is found, and true otherwise.
func (c *Cache[K, V]) Get(key K) (V, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    value, found := c.items[key]
    return value, found
}

// Remove deletes the key-value pair with the specified key from the cache.
func (c *Cache[K, V]) Remove(key K) {
    c.mu.Lock()
    defer c.mu.Unlock()

    delete(c.items, key)
}

// Pop removes and returns the value associated with the specified key from the cache.
func (c *Cache[K, V]) Pop(key K) (V, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    value, found := c.items[key]

    // If the key is found, delete the key-value pair from the cache.
    if found {
        delete(c.items, key)
    }

    return value, found
}
```

And you can use it like this:

```go
package main

import (
	"fmt"
	"time"

	"path/to/cache"
)

func main() {
	// Create a new Cache instance
	myCache := cache.New[string, int]()

	// Set key-value pairs in the cache
	myCache.Set("one", 1)
	myCache.Set("two", 2)
	myCache.Set("three", 3)

	// Retrieve values from the cache
	value, found := myCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache")
	}

	// Pop a key from the cache
	poppedValue, found := myCache.Pop("three")
	if found {
		fmt.Printf("Popped value for key 'three': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'three' not found in the cache")
	}

    // Remove a key from the cache
	myCache.Remove("one")

	// Try to retrieve a removed key
	removedValue, found := myCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", removedValue)
	} else {
		fmt.Println("Key 'one' not found in the cache (after removal)")
	}
}
```

## Expiring cache

You can extend this idea to associate an expiry time with every value in the cache, and launch a background goroutine to periodically remove expired entries. Like so:

```go
package cache

import (
    "sync"
    "time"
)

// item represents a cache item with a value and an expiration time.
type item[V any] struct {
    value  V           
    expiry time.Time  
}

// isExpired checks if the cache item has expired.
func (i item[V]) isExpired() bool {
    return time.Now().After(i.expiry)
}

// TTLCache is a generic cache implementation with support for time-to-live
// (TTL) expiration.
type TTLCache[K comparable, V any] struct {
    items map[K]item[V] // The map storing cache items.
    mu    sync.Mutex    // Mutex for controlling concurrent access to the cache.
}

// NewTTL creates a new TTLCache instance and starts a goroutine to periodically
// remove expired items every 5 seconds.
func NewTTL[K comparable, V any]() *TTLCache[K, V] {
    c := &TTLCache[K, V]{
        items: make(map[K]item[V]),
    }

    go func() {
        for range time.Tick(5 * time.Second) {
            c.mu.Lock()

            // Iterate over the cache items and delete expired ones.
            for key, item := range c.items {
                if item.isExpired() {
                    delete(c.items, key)
                }
            }

            c.mu.Unlock()
        }
    }()

    return c
}

// Set adds a new item to the cache with the specified key, value, and 
// time-to-live (TTL).
func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = item[V]{
        value:  value,
        expiry: time.Now().Add(ttl),
    }
}

// Get retrieves the value associated with the given key from the cache.
func (c *TTLCache[K, V]) Get(key K) (V, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    item, found := c.items[key]
    if !found {
        // If the key is not found, return the zero value for V and false.
        return item.value, false
    }

    if item.isExpired() {
        // If the item has expired, remove it from the cache and return the  
        // value and false.
        delete(c.items, key)
        return item.value, false
    }

    // Otherwise return the value and true.
    return item.value, true
}

// Remove removes the item with the specified key from the cache.
func (c *TTLCache[K, V]) Remove(key K) {
    c.mu.Lock()
    defer c.mu.Unlock()

    // Delete the item with the given key from the cache.
    delete(c.items, key)
}

// Pop removes and returns the item with the specified key from the cache.
func (c *TTLCache[K, V]) Pop(key K) (V, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    item, found := c.items[key]
    if !found {
        // If the key is not found, return the zero value for V and false.
        return item.value, false
    }

    // If the key is found, delete the item from the cache.
    delete(c.items, key)

    if item.isExpired() {
        // If the item has expired, return the value and false.
        return item.value, false
    }

    // Otherwise return the value and true.
    return item.value, true
}
```

And you can use this in much the same way:

```go
package main

import (
	"fmt"
	"time"

	"path/to/cache"
)

func main() {
	// Create a new TTLCache instance
	myTTLCache := cache.NewTTL[string, int]()

	// Set key-value pairs with TTL in the cache
	myTTLCache.Set("one", 1, 5*time.Second)
	myTTLCache.Set("two", 2, 10*time.Second)
	myTTLCache.Set("three", 3, 15*time.Second)

	// Retrieve values from the cache
	value, found := myTTLCache.Get("two")
	if found {
		fmt.Printf("Value for key 'two': %v\n", value)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Wait for a while to allow some items to expire
	time.Sleep(7 * time.Second)

	// Try to retrieve an expired key
	expiredValue, found := myTTLCache.Get("one")
	if found {
		fmt.Printf("Value for key 'one': %v\n", expiredValue)
	} else {
		fmt.Println("Key 'one' not found in the cache or has expired")
	}

	// Pop a key from the cache
	poppedValue, found := myTTLCache.Pop("two")
	if found {
		fmt.Printf("Popped value for key 'two': %v\n", poppedValue)
	} else {
		fmt.Println("Key 'two' not found in the cache or has expired")
	}

	// Remove a key from the cache
	myTTLCache.Remove("three")
}
```

