package gcitem

import "github.com/vprokopchuk256/gocache/pkg/gocmap/gclockable"

type Base struct {
	gclockable.Base
}

func base() Base {
	return Base{gclockable.New()}
}
