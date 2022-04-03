package lock

import "sync"

//var eventsLocks = make(map[string]sync.RWMutex)

type RWMap struct {
	locks map[string]sync.RWMutex
}

type Map struct {
	locks map[string]sync.Mutex
}

func NewRWMap() *RWMap {
	return &RWMap{
		locks: make(map[string]sync.RWMutex),
	}
}

func NewMap() *Map {
	return &Map{
		locks: make(map[string]sync.Mutex),
	}
}

func (t RWMap) Lock(name string) {
	if _, ok := t.locks[name]; !ok {
		t.locks[name] = sync.RWMutex{}
	}
	lock := t.locks[name]
	lock.Lock()
}

func (t RWMap) RLock(name string) {
	if _, ok := t.locks[name]; !ok {
		t.locks[name] = sync.RWMutex{}
	}
	lock := t.locks[name]
	lock.RLock()
}

func (t RWMap) Unlock(name string) {
	if lock, ok := t.locks[name]; ok {
		lock.Unlock()
	}
}

func (t RWMap) RUnlock(name string) {
	if lock, ok := t.locks[name]; ok {
		lock.RUnlock()
	}
}

func (t Map) Lock(name string) {
	if _, ok := t.locks[name]; !ok {
		t.locks[name] = sync.Mutex{}
	}
	lock := t.locks[name]
	lock.Lock()
}

func (t Map) Unlock(name string) {
	if lock, ok := t.locks[name]; ok {
		lock.Unlock()
	}
}
