package gcitem

import (
	"fmt"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gclockable"
)

type Item interface {
	gclockable.Lockable
	ExpireIn() int
	Expired() bool
	Expirable() bool
	SetExpireIn(int)
	Value() string
	fmt.Stringer
}
