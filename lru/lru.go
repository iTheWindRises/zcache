package lru

import "container/list"

type Cache struct {
	maxBytes int	// 运行使用最大内存
	nBytes int		// 当前使用内存
	ll *list.List
	cache map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key string
	value Value
}

type Value interface {
	Len() int
}



func New(maxBytes int, onEvicted func(string,Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value,true
	}
	return
}

// 淘汰
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache,kv.key)
		c.nBytes -= len(kv.key) +  kv.value.Len()
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
	return
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += value.Len() - kv.value.Len()
		kv.value = value
	}else {
		ele := c.ll.PushFront(&entry{key: key,value: value})
		c.cache[key] = ele
		c.nBytes += len(key) + value.Len()
	}
	// 当前使用内存超过最大内存淘汰
	for c.nBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}