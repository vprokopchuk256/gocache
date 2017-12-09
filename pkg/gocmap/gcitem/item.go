package gcitem

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gclockable"
)

type Item interface {
	gclockable.Lockable
	Value() string
	fmt.Stringer
}
