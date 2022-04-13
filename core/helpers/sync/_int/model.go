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
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = v.value + value
}

func (v *Int) Dec(value int) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = v.value - value
}

func (v *Int) Set(value int) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = value
}

func (v *Int) Get() int {
	v.lock.RLock()
	defer v.lock.RUnlock()
	return v.value
}
