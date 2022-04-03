package _struct

import "reflect"

func GetType(obj interface{}) string {
	return New(obj).GetType()
}
func (h *Helper) GetType() string {
	if t := reflect.TypeOf(h.NonPtrObj); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func GetName(obj interface{}) string {
	return GetType(obj)
}
