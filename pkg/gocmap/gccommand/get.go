package gccommand

import (
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Get(key string) (Command, error) {
	cmd := func(m *gcmap.Map) (string, error) {
		m.RLock()
		defer m.RUnlock()

		i, _ := m.Get(key)

		return i.Value(), nil
	}

	return cmd, nil
}
