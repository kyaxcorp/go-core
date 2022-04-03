package _map_string_interface

import (
	"sync"

	"github.com/kyaxcorp/go-core/core/helpers/function"
	"github.com/kyaxcorp/go-core/core/logger/appLog"
)

type MapStringInterface struct {
	// The lock for multiple goroutines to be able to access securely
	lock sync.RWMutex
	// Here we store the data!
	value map[string]interface{}
}

type Scan func(k string, v interface{})

func New() *MapStringInterface {
	return &MapStringInterface{
		lock:  sync.RWMutex{},
		value: make(map[string]interface{}),
	}
}

func (v *MapStringInterface) Set(key string, value interface{}) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.value[key] = value
}

func (v *MapStringInterface) Del(key string) {
	defer v.lock.Unlock()
	v.lock.Lock()
	if _, ok := v.value[key]; ok {
		delete(v.value, key)
	}
}

func (v *MapStringInterface) Has(key string) bool {
	defer v.lock.RUnlock()
	v.lock.RLock()
	if _, ok := v.value[key]; ok {
		return true
	}
	return false
}

func (v *MapStringInterface) Get(key string) interface{} {
	defer v.lock.RUnlock()
	v.lock.RLock()
	if vv, ok := v.value[key]; ok {
		return vv
	}
	return nil
}

func (v *MapStringInterface) Len() int {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return len(v.value)
}

func (v *MapStringInterface) Scan(callback Scan) {
	// TODO: we should add here panic recover?!
	if v.Len() == 0 {
		return
	}
	if !function.IsCallable(callback) {
		return
	}
	defer func() {
		v.lock.RUnlock()
		// Recover here from panicks!
		if r := recover(); r != nil {
			appLog.Warn().Interface("recover_stack", r).Msg("Scan -> Recovered from panic")
		}
	}()
	v.lock.RLock()
	for k, vv := range v.value {
		// Callbacks can have panicks, so we added recover!
		callback(k, vv)
	}
}
