package sync

import (
	"reflect"
	"sync"
)

type RWMutex struct {
	// This is the lock itself
	lock sync.RWMutex
	// this is the tryLock,tryReadLock,tryWriteLock operation lock -> it will help as for multiple routines to not block each other!
	tryLock      sync.Mutex
	tryReadLock  sync.Mutex
	tryWriteLock sync.Mutex
}

// NewRW -> It's not mandatory to use, because when defining the variable with type Mutex, it's automatically defined with
// zero values!
func NewRW() *RWMutex {
	return &RWMutex{
		lock:    sync.RWMutex{},
		tryLock: sync.Mutex{},
	}
}

func (l *RWMutex) TryLock() bool {
	// Check if it's not locked!
	// If it's not then lock!
	// If it's locked, return!
	// There is a problem here, concurrent routines can call same function and bypass the IsLocked function, and on
	// calling lock function, 1 routine will lock, but the other one it will be blocked and wait until the first
	// one is release
	// A solution is to create another lock here, which will block the operation of locking itself! after the lock is
	// acquired, it will release the operation lock!
	l.tryLock.Lock()
	defer l.tryLock.Unlock()
	if l.IsLocked() {
		return false
	}
	l.lock.Lock()
	return true
}

func (l *RWMutex) TryReadLock() bool {
	l.tryReadLock.Lock()
	defer l.tryReadLock.Unlock()
	if l.IsReadLocked() {
		return false
	}
	l.lock.RLock()
	return true
}

func (l *RWMutex) Lock() {
	l.lock.Lock()
}

func (l *RWMutex) RLock() {
	l.lock.RLock()
}

func (l *RWMutex) Unlock() {
	l.lock.Unlock()
}

func (l *RWMutex) RUnlock() {
	l.lock.RUnlock()
}

func (l *RWMutex) IsLocked() bool {
	state := reflect.ValueOf(&l.lock).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

func (l *RWMutex) IsWriteLocked() bool {
	// RWMutex has a "w" sync.Mutex field for write lock
	state := reflect.ValueOf(&l.lock).Elem().FieldByName("w").FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}

func (l *RWMutex) IsReadLocked() bool {
	return reflect.ValueOf(&l.lock).Elem().FieldByName("readerCount").Int() > 0
}
