package gclockable

import "sync"

type Base struct {
	lock *sync.RWMutex
}

func New() Base {
	return Base{lock: new(sync.RWMutex)}
}

func (l *Base) Lock() {
	l.lock.Lock()
}

func (l *Base) Unlock() {
	l.lock.Unlock()
}

func (l *Base) RLock() {
	l.lock.RLock()
}

func (l *Base) RUnlock() {
	l.lock.RUnlock()
}
