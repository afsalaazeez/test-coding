package myfunc

import (
	"container/list"
	"fmt"
)

// OMap structure defines ordered map structure
type OMap struct {
	l *list.List
	m map[interface{}]*list.Element
}

// OMapConstr define ordered map constructor
func OMapConstr() *OMap {
	m := new(OMap)
	m.l = list.New()
	m.m = make(map[interface{}]*list.Element)
	return m
}

// Add function to add elements
func (m *OMap) Add(key interface{}, val interface{}) {
	e := m.l.PushBack(val)
	m.m[key] = e
}

// Show function to show elements in structure
func (m *OMap) Show() {
	for e := m.l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// Delete function to delete elements
func (m *OMap) Delete(key interface{}) {
	if node, ok := m.m[key]; ok {
		m.l.Remove(node)
		delete(m.m, key)
	}
}

// Len function returns number of elements in the ordered map
func (m *OMap) Len() interface{} {
	return m.l.Len()
}

// Get function returns value of the key parsed
func (m *OMap) Get(key interface{}) (interface{}, bool) {
	if node, ok := m.m[key]; ok {
		return node.Value, ok
	}
	return 0, false
}
