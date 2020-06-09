package hw04_lru_cache //nolint:golint,stylecheck
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
	items    map[Key]*listItem
	mux      sync.Mutex
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	element, exists := c.items[key]

	if exists {
		c.queue.MoveToFront(element)
		element.Value.(*cacheItem).Value = value

		return true
	}

	if c.queue.Len() == c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		delete(c.items, lastItem.Value.(*cacheItem).Key)
	}

	item := &cacheItem{
		Key:   key,
		Value: value,
	}

	element = c.queue.PushFront(item)
	c.items[item.Key] = element

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	element, exists := c.items[key]

	if exists {
		c.queue.MoveToFront(element)

		return element.Value.(*cacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.items = make(map[Key]*listItem)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		items:    make(map[Key]*listItem),
		queue:    NewList(),
	}
}
