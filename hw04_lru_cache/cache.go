package hw04lrucache

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
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	ci := cacheItem{key, value}
	if li, found := c.items[key]; found {
		li.Value = value
		c.queue.MoveToFront(li)
		return true
	}
	qp := c.queue.PushFront(ci.value)
	c.items[key] = qp
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if li, found := c.items[key]; found {
		c.queue.MoveToFront(li)
		return li.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
