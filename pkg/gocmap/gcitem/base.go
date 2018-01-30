package gcitem

import (
	"time"

	"github.com/vprokopchuk256/gocache/pkg/gocmap/gclockable"
)

type Base struct {
	gclockable.Base
	expireAt time.Time
}

func base() Base {
	return Base{gclockable.New(), time.Time{}}
}

func (e *Base) ExpireIn() int {
	if e.Expirable() {
		return int(time.Until(e.expireAt).Seconds())
	}

	return int(^uint(0) >> 1)
}

func (e *Base) Expired() bool {
	return e.ExpireIn() <= 0
}

func (e *Base) Expirable() bool {
	return !e.expireAt.IsZero()
}

func (e *Base) SetExpireIn(expireIn int) {
	e.expireAt = expireAt(expireIn)
}

func expireAt(expireIn int) time.Time {
	return time.Now().Local().Add(time.Second * time.Duration(expireIn))
}
