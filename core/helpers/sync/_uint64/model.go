package _uint64

import "sync"

type Uint64 struct {
	// TODO: Should we use Atomic instead of sync?
	lock  sync.RWMutex
	value uint64
}

func New() *Uint64 {
	return &Uint64{
		lock:  sync.RWMutex{},
		value: 0,
	}
}

func NewVal(v uint64) *Uint64 {
	return &Uint64{
		lock:  sync.RWMutex{},
		value: v,
	}
}

func (v *Uint64) Inc(value uint64) uint64 {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value + value
	return v.value
}

// ResetMaxAndInc -> increment by a specific value and set a Reset Maximum when it should reset!
// set the specified value, and after that increment it!
func (v *Uint64) ResetMaxAndInc(
	incVal uint64,
	resetMax uint64,
	resetTo uint64,
) uint64 {
	defer v.lock.Unlock()
	v.lock.Lock()
	if v.value >= resetMax {
		v.value = resetTo
	}
	v.value = v.value + incVal
	return v.value
}

func (v *Uint64) Dec(value uint64) uint64 {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = v.value - value
	return v.value
}

func (v *Uint64) Set(value uint64) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *Uint64) Get() uint64 {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
