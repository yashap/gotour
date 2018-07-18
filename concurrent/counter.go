package concurrent

import (
	"sync"
)

// Counter is safe to use concurrently
type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

// NewCounter constructs a Counter
func NewCounter() Counter {
	return Counter{v: make(map[string]int)}
}

// Inc increments the count for a given key
func (c *Counter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter at the given key
func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
