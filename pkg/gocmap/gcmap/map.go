package gcmap

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gclockable"
)

type Map struct {
	gclockable.Base
	mp map[string]gcitem.Item
}

func New() *Map {
	return &Map{gclockable.New(), make(map[string]gcitem.Item)}
}

func (m *Map) Set(key string, i gcitem.Item) {
	m.mp[key] = i
}

func (m *Map) Get(key string) (i gcitem.Item, ok bool) {
	item, ok := m.mp[key]

	return item, ok
}

func (m *Map) Delete(key string) {
	delete(m.mp, key)
}
