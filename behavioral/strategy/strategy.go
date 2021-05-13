package strategy

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Strategy_pattern

/**
 * Strategy is a behavioral design pattern that define a family of algorithms, put each of them
 * into a separate class, and make their objects interchangeable.
 */

// strategy interface - evictionAlgo
type evictionAlgo interface {
	evict(c *cache)
}

// concrete strategy - fifo
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// concrete strategy - lru
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

// concrete strategy - lfu
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

// context - cache
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
