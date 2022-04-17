package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       *sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       new(sync.Mutex),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()

	item := &cacheItem{key: key, value: value}
	if i, ok := c.items[key]; ok {
		i.Value = item
		c.queue.MoveToFront(i)

		c.mu.Unlock()
		return true
	}

	pushed := c.queue.PushFront(item)

	if c.queue.Len() > c.capacity {
		lastItem, ok := c.queue.Back().Value.(*cacheItem)
		if !ok {
			return false
		}

		c.queue.Remove(c.queue.Back())
		delete(c.items, lastItem.key)
	}

	c.items[key] = pushed
	c.mu.Unlock()

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		cachedItem, ok := item.Value.(*cacheItem)
		if !ok {
			return nil, false
		}

		c.mu.Unlock()
		return cachedItem.value, true
	}

	c.mu.Unlock()
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	for i := range c.items {
		delete(c.items, i)
	}
}
