package cache

import (
	"fmt"
	"sync"
)

type Item interface{}

type Memorycache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func NewCache() *Memorycache {
	items := make(map[string]Item)

	c := &Memorycache{
		items: items,
	}

	return c
}

func (c *Memorycache) Invalidate() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}

func (c *Memorycache) Add(key string, item Item) error {
	c.mu.Lock()
	_, found := c.items[key]
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s already exists", key)
	}
	c.items[key] = item
	c.mu.Unlock()
	return nil
}

func (c *Memorycache) Get(key string) (Item, bool) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}

	c.mu.RUnlock()
	return item, true
}
