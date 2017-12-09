package gccommand

import "github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"

type Command interface {
	Exec(m *gcmap.Map) (string, error)
}
