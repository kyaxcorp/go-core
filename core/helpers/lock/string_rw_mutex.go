package lock

import "sync"

// TODO: COpy the functionality from the string_mutex.go

type RWString struct {
	lock  sync.Mutex
	locks map[string]*sync.RWMutex
}

func NewRWString() *RWString {
	return &RWString{
		lock:  sync.Mutex{},
		locks: make(map[string]*sync.RWMutex),
	}
}

func (t *RWString) Lock(name string) {
	t.lock.Lock()

	var l *sync.RWMutex
	var ok bool

	if l, ok = t.locks[name]; !ok {
		l = &sync.RWMutex{}
		// Write into the Map the lock
		t.locks[name] = l
	}
	// Unlock the map lock
	t.lock.Unlock()
	// call the string lock!
	l.Lock()
}

func (t *RWString) RLock(name string) {
	t.lock.Lock()

	var l *sync.RWMutex
	var ok bool

	if l, ok = t.locks[name]; !ok {
		l = &sync.RWMutex{}
		// Write into the Map the lock
		t.locks[name] = l
	}
	// Unlock the map lock
	t.lock.Unlock()
	// call the string lock!
	l.RLock()
}

func (t *RWString) Unlock(name string) {
	t.lock.Lock()

	var l *sync.RWMutex
	var ok bool

	if l, ok = t.locks[name]; ok {
		// Unlock the map lock
		// TODO: be careful when we will be creating the TTL and garbage collector
		t.lock.Unlock()
		// call the string Unlock!
		l.Unlock()
	}
}

func (t *RWString) RUnlock(name string) {
	t.lock.Lock()

	var l *sync.RWMutex
	var ok bool

	if l, ok = t.locks[name]; ok {
		// Unlock the map lock
		// TODO: be careful when we will be creating the TTL and garbage collector
		t.lock.Unlock()
		// call the string Unlock!
		l.RUnlock()
	}
}
