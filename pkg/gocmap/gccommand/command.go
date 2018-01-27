package gccommand

import "github.com/vprokopchuk256/gocache/pkg/gocmap/gcmap"

type Command func(m *gcmap.Map) (string, error)
