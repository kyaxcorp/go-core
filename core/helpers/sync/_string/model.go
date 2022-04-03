package _string

import "sync"

type String struct {
	lock  sync.RWMutex
	value string
}

func New() *String {
	return &String{
		lock:  sync.RWMutex{},
		value: "",
	}
}

func (v *String) Set(value string) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value = value
}

func (v *String) Get() string {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}
