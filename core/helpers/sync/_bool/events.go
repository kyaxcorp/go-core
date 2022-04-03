package _bool

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/function"
)

type OnChange func(v *Bool, newVal bool)
type OnChangeAsync func(v *Bool, newVal bool)

// OnTrue -> when True happens!
func (v *Bool) OnTrue(eventName string) {

}

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
