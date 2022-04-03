package _uint16

import "sync"

type Uint16 struct {
	// TODO: Should we use Atomic instead of sync?
	lock  sync.RWMutex
	value uint16
}

func New() *Uint16 {
	return &Uint16{
		lock:  sync.RWMutex{},
		value: 0,
	}
}

func NewVal(v uint16) *Uint16 {
	return &Uint16{
		lock:  sync.RWMutex{},
		value: v,
	}
}

func (v *Uint16) Inc(value uint16) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value + value
}

func (v *Uint16) Dec(value uint16) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value - value
}

func (v *Uint16) Set(value uint16) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *Uint16) Get() uint16 {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
