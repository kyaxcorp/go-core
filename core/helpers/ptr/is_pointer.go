package ptr

import "reflect"

func Is(v interface{}) bool {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		return true
	}
	return false
}
