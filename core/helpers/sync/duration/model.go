package duration

import (
	"sync"
	"time"
)

type Duration struct {
	lock  sync.RWMutex
	value time.Duration
}

func New() *Duration {
	return &Duration{
		lock:  sync.RWMutex{},
		value: 0,
	}
}

func NewVal(t time.Duration) *Duration {
	return &Duration{
		lock:  sync.RWMutex{},
		value: t,
	}
}

func (v *Duration) Set(value time.Duration) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *Duration) Get() time.Duration {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
