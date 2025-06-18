package kit

import "sync"

func WithLock(l sync.Locker, fn func()) {
	l.Lock()
	defer l.Unlock()
	fn()
}
