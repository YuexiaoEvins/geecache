package lru

import "container/list"

type Cache struct {
	MaxEntries int
	curEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
}

type cacheEntry struct {
	key   interface{}
	value interface{}
}

func NewLruCache(max int) *Cache {
	return &Cache{
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element, 0),
		MaxEntries: max,
	}
}

func (c *Cache) Set(key, value interface{}) {
	if elem, ok := c.cache[key]; ok {
		//if exist and overwrite
		c.ll.MoveToFront(elem)
		kv := elem.Value.(cacheEntry)
		kv.value = value
	} else {
		//if non-exist
		preElem := c.ll.PushFront(&cacheEntry{key: key, value: value})
		c.cache[key] = preElem
		c.curEntries += 1
		if c.MaxEntries != 0 && c.curEntries > c.MaxEntries {
			c.RemoveOldest()
		}
	}
}

func (c *Cache) RemoveOldest() {
	elem := c.ll.Back()
	if elem != nil {
		c.ll.Remove(elem)
		delete(c.cache, elem.Value.(cacheEntry).key)
		c.curEntries -= 1
	}
}

func (c *Cache) Get(key string) interface{} {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*cacheEntry).value
	}
	return nil
}

func (c *Cache) Debug() {

}

func (c *Cache) Len() int {
	return c.ll.Len()
}
