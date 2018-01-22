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
	return Base{gclockable.New(), expireAt(1000)}
}

func (e *Base) ExpireIn() int {
	return int(time.Until(e.expireAt).Seconds())
}

func (e *Base) Expired() bool {
	return e.ExpireIn() <= 0
}

func (e *Base) SetExpireIn(expireIn int) {
	e.expireAt = expireAt(expireIn)
}

func expireAt(expireIn int) time.Time {
	return time.Now().Local().Add(time.Second * time.Duration(expireIn))
}
