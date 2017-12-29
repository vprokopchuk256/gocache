package gccommand

import (
	"fmt"
	"strings"

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

func NewInsert(key string, value string) (*Insert, error) {
	i, err := gcitem.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("could not create insert operation: %v", err)
	}

	return &Insert{key: key, item: i}, nil
}

func ParseInsert(params string) (Command, error) {
	parts := strings.SplitAfterN(params, " ", 2)

	if len(parts) != 2 {
		return nil, fmt.Errorf("could not recognize key and params for insert: %v", params)
	}

	return NewInsert(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
}
