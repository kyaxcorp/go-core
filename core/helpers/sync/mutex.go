package sync

import (
	"reflect"
	"sync"
)

const mutexLocked = 1

type Mutex struct {
	// This is the lock itself
	lock sync.Mutex
	// this is the tryLock operation lock -> it will help as for multiple routines to not block each other!
	tryLock sync.Mutex
}

// New -> It's not mandatory to use, because when defining the variable with type Mutex, it's automatically defined with
// zero values!
func New() *Mutex {
	return &Mutex{
		lock:    sync.Mutex{},
		tryLock: sync.Mutex{},
	}
}

func (l *Mutex) TryLock() bool {
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

func (l *Mutex) Lock() {
	l.lock.Lock()
}

func (l *Mutex) Unlock() {
	l.lock.Unlock()
}

func (l *Mutex) IsLocked() bool {
	state := reflect.ValueOf(&l.lock).Elem().FieldByName("state")
	return state.Int()&mutexLocked == mutexLocked
}
