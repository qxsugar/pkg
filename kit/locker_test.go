package kit

import (
	"sync"
	"testing"
)

type MockLocker struct {
	mutex        sync.Mutex
	lockCalled   bool
	unlockCalled bool
}

func (m *MockLocker) Lock() {
	m.lockCalled = true
	m.mutex.Lock()
}

func (m *MockLocker) Unlock() {
	m.unlockCalled = true
	m.mutex.Unlock()
}

func TestWithLock(t *testing.T) {
	mockLocker := &MockLocker{}
	isFunctionCalled := false

	WithLock(mockLocker, func() {
		if !mockLocker.lockCalled {
			t.Error("Lock was not called before function execution")
		}

		isFunctionCalled = true
	})

	if !isFunctionCalled {
		t.Error("The provided function was not called")
	}

	if !mockLocker.unlockCalled {
		t.Error("Unlock was not called after function execution")
	}
}
