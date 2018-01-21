package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

type Get struct {
	key string
}

func NewGet(key string) (*Get, error) {
	return &Get{key: key}, nil
}

func ParseGet(key string) (Command, error) {
	return NewGet(key)
}

func (c *Get) Key() string {
	return c.key
}

func (c *Get) Exec(m *gcmap.Map) (string, error) {
	m.RLock()
	defer m.RUnlock()

	i, _ := m.Get(c.key)

	return fmt.Sprintf("%v := %v", c.key, i), nil
}
