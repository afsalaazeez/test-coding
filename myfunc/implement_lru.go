package myfunc

import (
	"container/list"
	"fmt"
)

// LRUCache data structure
type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

// Pair is value of Element
type Pair struct {
	key   int
	value int
}

// LRUConstructor initializes a LRUCache
func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get item considered as item access/use for LRU to be updated
func (c *LRUCache) Get(key int) interface{} {
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		c.l.MoveToFront(node)
		return val
	}
	return "miss"
}

// Len function
func (c *LRUCache) Len() int {
	return (c.l.Len())
}

// Show function
func (c *LRUCache) Show() {
	fmt.Printf("Cache Values: ")
	for e := c.l.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %d", e.Value.(*list.Element).Value.(Pair).value)
	}
	fmt.Println()
}

// Add key and value in the LRUCache
func (c *LRUCache) Add(key int, value int) {
	if node, ok := c.m[key]; ok {
		c.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		if c.l.Len() == c.cap {
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			delete(c.m, idx)
			c.l.Remove(c.l.Back())
		}
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := c.l.PushFront(node)
		c.m[key] = ptr
	}
}
