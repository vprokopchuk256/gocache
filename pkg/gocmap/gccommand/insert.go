package gccommand

import (
	"fmt"
	"strings"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcitem"
	"github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"
)

func Insert(params string) (Command, error) {
	parseCommand := func() (string, string, error) {
		parts := strings.SplitAfterN(params, " ", 2)

		if len(parts) != 2 {
			return "", "", fmt.Errorf("could not recognize key and params for insert: %v", params)
		}

		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
	}

	key, value, err := parseCommand()
	if err != nil {
		return nil, err
	}

	item, err := gcitem.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("could not create insert operation: %v", err)
	}

	cmd := func(m *gcmap.Map) (string, error) {
		m.Lock()
		defer m.Unlock()

		m.Set(key, item)

		return fmt.Sprintf("%v := %v", key, item), nil
	}

	return cmd, nil
}
