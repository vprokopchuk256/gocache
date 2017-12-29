package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

type Inc struct {
	key string
}

func NewInc(key string) (*Inc, error) {
	return &Inc{key: key}, nil
}

func ParseInc(key string) (Command, error) {
	return NewInc(key)
}

func (c *Inc) Key() string {
	return c.key
}

func (c *Inc) Exec(m *gcmap.Map) (string, error) {
	m.RLock()
	defer m.RUnlock()

	i, _ := m.Get(c.key)

	switch i.(type) {
	case *gcitem.Integer:
		return fmt.Sprintf("%v := %v", c.key, inc(i)), nil
	default:
		return "", fmt.Errorf("Value %v does not support inc operation", i)

	}
}

func inc(i gcitem.Item) gcitem.Item {
	return i.(*gcitem.Integer).Inc()
}
