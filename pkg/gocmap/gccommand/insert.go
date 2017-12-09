package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

type Insert struct {
	key  string
	item gcitem.Item
}

func (c *Insert) Key() string {
	return c.key
}

func (c *Insert) Item() gcitem.Item {
	return c.item
}

func (c *Insert) Exec(m *gcmap.Map) (string, error) {
	m.Lock()
	defer m.Unlock()

	m.Set(c.key, c.item)

	return fmt.Sprintf("%v := %v", c.key, c.item), nil
}

func NewInsert(key string, value string) (*Insert, bool) {
	i, ok := gcitem.Parse(value)
	if !ok {
		return nil, false
	}

	return &Insert{key: key, item: i}, true
}
