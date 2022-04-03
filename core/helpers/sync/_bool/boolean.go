package _bool

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

func (v *Bool) Set(value bool) {
	defer v.lock.Unlock()
	v.lock.Lock()
	if value != v.value {
		v.value = value
		// Trigger Change

		if value {
			v.onTrueWaiter.Signal()
		} else {
			v.onFalseWaiter.Signal()
		}

		// Launch async
		go func(value bool) {
			v.onChangeAsync.Scan(func(k string, vv interface{}) {
				vv.(OnChange)(v, value)
			})
		}(value)

		// Execute in current goroutine
		v.onChange.Scan(func(k string, vv interface{}) {
			vv.(OnChange)(v, value)
		})
	}
}

func (v *Bool) True() {
	// Check if already true
	if v.Get() == true {
		return
	}
	v.Set(true)
}

func (v *Bool) WaitForTrue() *Bool {
	return v.WaitUntilTrue()
}

func (v *Bool) WaitForFalse() *Bool {
	return v.WaitUntilFalse()
}

// WaitUntilTrue -> If it's not true, it will wait until is it!
// If it's true already, it will go futher without waiting!
func (v *Bool) WaitUntilTrue() *Bool {
	// Create the channel
	if !v.Get() {
		v.onTrueWaiter.Wait()
	}
	return v
}

// WaitUntilFalse -> If it's not false, it will wait until is it!
// If it's false already, it will go futher without waiting!
func (v *Bool) WaitUntilFalse() *Bool {
	// if it's true, wait...
	if v.Get() {
		v.onFalseWaiter.Wait()
	}
	return v
}

func (v *Bool) False() {
	// Check if already false
	if v.Get() == false {
		return
	}
	v.Set(false)
}

func (v *Bool) Get() bool {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.value
}

func (v *Bool) GetAndToggle() bool {
	defer v.lock.Unlock()
	v.lock.Lock()
	current := v.value
	v.value = !v.value
	return current
}

func (v *Bool) IfTrueSetFalse() bool {
	defer v.lock.Unlock()
	v.lock.Lock()
	current := v.value
	if v.value {
		v.value = false
	}
	return current
}

func (v *Bool) IfFalseSetTrue() bool {
	defer v.lock.Unlock()
	v.lock.Lock()
	current := v.value
	if !v.value {
		v.value = true
	}
	return current
}

func (v *Bool) RWGet(callback func(v *Bool)) bool {
	defer v.lock.Unlock()
	if function.IsCallable(callback) {
		callback(v)
	}
	v.lock.Lock()
	return v.value
}
