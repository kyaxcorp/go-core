package _bool

import (
	"github.com/kyaxcorp/go-core/core/helpers/function"
)

type OnChange func(v *Bool, newVal bool)
type OnTrue func(v *Bool)
type OnFalse func(v *Bool)
type OnChangeAsync func(v *Bool, newVal bool)

//================== On True==================\\

// OnTrue -> when True happens!
func (v *Bool) OnTrue(eventName string, onTrue OnTrue) bool {
	if !function.IsCallable(onTrue) || eventName == "" {
		return false
	}
	v.onTrue.Set(eventName, onTrue)
	return true
}

func (v *Bool) OnTrueRemove(eventName string) *Bool {
	v.onTrue.Del(eventName)
	return v
}

func (v *Bool) HasOnTrue(eventName string) bool {
	return v.onTrue.Has(eventName)
}

//================== On True==================\\

//

//================== On True Async ==================\\

// OnTrue -> when True happens!
func (v *Bool) OnTrueAsync(eventName string, onTrue OnTrue) bool {
	if !function.IsCallable(onTrue) || eventName == "" {
		return false
	}
	v.onTrueAsync.Set(eventName, onTrue)
	return true
}

func (v *Bool) OnTrueAsyncRemove(eventName string) *Bool {
	v.onTrueAsync.Del(eventName)
	return v
}

func (v *Bool) HasOnTrueAsync(eventName string) bool {
	return v.onTrueAsync.Has(eventName)
}

//================== On True Async ==================\\

//

//=================== On False ==================== \\

// OnFalse -> when false happens!
func (v *Bool) OnFalse(eventName string, onFalse OnFalse) bool {
	if !function.IsCallable(onFalse) || eventName == "" {
		return false
	}
	v.onFalse.Set(eventName, onFalse)
	return true
}

func (v *Bool) OnFalseRemove(eventName string) *Bool {
	v.onFalse.Del(eventName)
	return v
}

func (v *Bool) HasOnFalse(eventName string) bool {
	return v.onFalse.Has(eventName)
}

//=================== On False ==================== \\

//

//=================== On False Async ==================== \\

// OnFalse -> when false happens!
func (v *Bool) OnFalseAsync(eventName string, onFalse OnFalse) bool {
	if !function.IsCallable(onFalse) || eventName == "" {
		return false
	}
	v.onFalseAsync.Set(eventName, onFalse)
	return true
}

func (v *Bool) OnFalseAsyncRemove(eventName string) *Bool {
	v.onFalseAsync.Del(eventName)
	return v
}

func (v *Bool) HasOnFalseAsync(eventName string) bool {
	return v.onFalseAsync.Has(eventName)
}

//=================== On False Async ==================== \\

//

//

//--------------------ON CHANGE----------------------\\

// OnChange -> when the value has changed!
func (v *Bool) OnChange(eventName string, onChange OnChange) bool {
	if !function.IsCallable(onChange) || eventName == "" {
		return false
	}
	v.onChange.Set(eventName, onChange)
	return true
}

func (v *Bool) OnChangeRemove(eventName string) *Bool {
	v.onChange.Del(eventName)
	return v
}

func (v *Bool) HasOnChange(eventName string) bool {
	return v.onChange.Has(eventName)
}

//--------------------ON CHANGE----------------------\\

//

//--------------------ON CHANGE ASYNC----------------------\\

// OnChangeAsync -> when the value has changed!
func (v *Bool) OnChangeAsync(eventName string, onChange OnChange) bool {
	if !function.IsCallable(onChange) || eventName == "" {
		return false
	}
	v.onChangeAsync.Set(eventName, onChange)
	return true
}

func (v *Bool) OnChangeAsyncRemove(eventName string) *Bool {
	v.onChangeAsync.Del(eventName)
	return v
}

func (v *Bool) HasOnChangeAsync(eventName string) bool {
	return v.onChangeAsync.Has(eventName)
}

//--------------------ON CHANGE ASYNC----------------------\\

//
