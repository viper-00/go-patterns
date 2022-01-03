package strategy

import "fmt"

// wiki: https://en.wikipedia.org/wiki/Strategy_pattern
//
// Enables an algorithm's behavior to be selected at runtime.

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
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

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
