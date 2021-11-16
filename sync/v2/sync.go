package v2

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}
func (c *Counter) Inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}
