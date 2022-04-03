package _time

import (
	"sync"
	"time"
)

type Time struct {
	lock  sync.RWMutex
	value time.Time
}

func New() *Time {
	return &Time{
		lock:  sync.RWMutex{},
		value: time.Time{},
	}
}

func NewNow() *Time {
	return NewVal(time.Now())
}

func NewVal(t time.Time) *Time {
	return &Time{
		lock:  sync.RWMutex{},
		value: t,
	}
}

func (v *Time) Set(value time.Time) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *Time) SetNow() {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = time.Now()
}

func (v *Time) Get() time.Time {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
