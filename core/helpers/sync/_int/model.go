package _int

import "sync"

type Int struct {
	// TODO: Should we use Atomic instead of sync?
	lock  sync.RWMutex
	value int
}

func New() *Int {
	return &Int{
		lock:  sync.RWMutex{},
		value: 0,
	}
}

func NewVal(v int) *Int {
	return &Int{
		lock:  sync.RWMutex{},
		value: v,
	}
}

func (v *Int) Inc(value int) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value + value
}

func (v *Int) Dec(value int) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value - value
}

func (v *Int) Set(value int) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *Int) Get() int {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
