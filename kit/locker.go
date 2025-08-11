package kit

import "sync"

// WithLock executes the given function while holding the provided lock.
// The lock is automatically released when the function completes.
func WithLock(l sync.Locker, fn func()) {
	l.Lock()
	defer l.Unlock()
	fn()
}
