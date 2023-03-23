package lru

import "container/list"

type Cache struct {
	nbytes    int64
	maxbytes  int64
	ll        *list.List
	cache     map[string]*list.Element
	onEvicted func(key string, value Value)
}
type entry struct {
	key string
	val Value
}

type Value interface {
	Len() int
}

func New(maxbytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{maxbytes: maxbytes, onEvicted: onEvicted, ll: list.New(), cache: make(map[string]*list.Element)}
}
func (c *Cache) Add(key string, val Value) {
	if ele, ok := c.cache[key]; ok {
		kv := ele.Value.(*entry)
		c.ll.MoveToFront(ele)
		kv.val = val
		c.nbytes += int64(val.Len()) - int64(kv.val.Len())
	} else {
		ele = c.ll.PushFront(&entry{key, val})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(val.Len())
	}
	for c.maxbytes != 0 && c.maxbytes < c.nbytes {
		c.Remove()
	}
}
func (c *Cache) Remove() {
	if ele := c.ll.Back(); ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.val.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.val)
		}
	}
}
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.val, true
	}
	return
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
