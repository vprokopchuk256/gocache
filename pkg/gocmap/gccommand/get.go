package gccommand

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Get(key string) (Command, error) {
	cmd := func(m *gcmap.Map) (string, error) {
		m.RLock()
		defer m.RUnlock()

		i, _ := m.Get(key)

		return fmt.Sprintf("%v := %v", key, i), nil
	}

	return cmd, nil
}
