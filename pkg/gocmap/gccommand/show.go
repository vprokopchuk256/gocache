package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Show(key string) (Command, error) {
	cmd := func(m *gcmap.Map) (string, error) {
		m.RLock()
		defer m.RUnlock()

		item, _ := m.Get(key)

		return fmt.Sprintf("%v := %v", key, item), nil
	}

	return cmd, nil
}
