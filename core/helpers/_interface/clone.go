package _interface

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_struct"
	"reflect"
)

// CloneInterfaceItem -> if you have a variable interface{}, it will take whatever value from there and
// it will clone to a new address so you can change the contents of it!
// it will return a new pointer/address of that value
func CloneInterfaceItem(obj interface{}) interface{} {
	var i interface{}
	if _struct.IsPointer(obj) {
		i = reflect.Indirect(reflect.ValueOf(obj)).Interface()
	} else {
		i = obj
	}

	p := reflect.New(reflect.TypeOf(i))
	p.Elem().Set(reflect.ValueOf(i))
	return p.Interface()
}
