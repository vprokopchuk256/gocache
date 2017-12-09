package gclockable

type Lockable interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}
