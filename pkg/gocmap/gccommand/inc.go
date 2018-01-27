package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Inc(key string) (Command, error) {
	inc := func(i gcitem.Item) gcitem.Item {
		return i.(*gcitem.Integer).Inc()
	}

	cmd := func(m *gcmap.Map) (string, error) {
		m.RLock()
		defer m.RUnlock()

		i, _ := m.Get(key)

		switch i.(type) {
		case *gcitem.Integer:
			return fmt.Sprintf("%v := %v", key, inc(i)), nil
		default:
			return "", fmt.Errorf("Value %v does not support inc operation", i)
		}
	}

	return cmd, nil
}
