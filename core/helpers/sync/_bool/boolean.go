package _bool

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

func (v *Bool) Set(value bool) {
	v.lock.Lock()
	defer v.lock.Unlock()
	if value != v.value {
		v.value = value
		// Trigger Change

		if value {
			v.onTrueWaiter.Signal()
		} else {
			v.onFalseWaiter.Signal()
		}

		// TODO: if we want all events to be non blocking we should
		// set an option here for that!

		// Launch async
		go func(value bool) {
			v.onChangeAsync.Scan(func(k string, vv interface{}) {
				go vv.(OnChange)(v, value)
			})
		}(value)

		// Execute in current goroutine
		v.onChange.Scan(func(k string, vv interface{}) {
			vv.(OnChange)(v, value)
		})

		if value {
			// if true
			go func(value bool) {
				v.onTrueAsync.Scan(func(k string, vv interface{}) {
					go vv.(OnTrue)(v)
				})
			}(value)

			v.onTrue.Scan(func(k string, vv interface{}) {
				vv.(OnTrue)(v)
			})
		} else {
			// if false
			go func(value bool) {
				v.onFalseAsync.Scan(func(k string, vv interface{}) {
					go vv.(OnFalse)(v)
				})
			}(value)

			v.onFalse.Scan(func(k string, vv interface{}) {
				vv.(OnFalse)(v)
			})
		}
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
	// if it's not true (it's false) then let's access the wait
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
	v.lock.RLock()
	defer v.lock.RUnlock()
	return v.value
}

func (v *Bool) GetAndToggle() bool {
	v.lock.Lock()
	defer v.lock.Unlock()
	current := v.value
	v.value = !v.value
	return current
}

func (v *Bool) IfTrueSetFalse() bool {
	v.lock.Lock()
	defer v.lock.Unlock()
	current := v.value
	if v.value {
		v.value = false
	}
	return current
}

func (v *Bool) IfFalseSetTrue() bool {
	v.lock.Lock()
	defer v.lock.Unlock()
	current := v.value
	if !v.value {
		v.value = true
	}
	return current
}

// TODO: something is wrong here?!!
func (v *Bool) RWGet(callback func(v *Bool)) bool {
	defer v.lock.Unlock()
	if function.IsCallable(callback) {
		callback(v)
	}
	v.lock.Lock()
	return v.value
}
