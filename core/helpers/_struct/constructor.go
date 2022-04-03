package _struct

import (
	"github.com/kyaxcorp/go-core/core/helpers/ptr"
	"reflect"
)

func New(obj interface{}) *Helper {
	if obj == nil {
		panic("obj is nil")
	}

	isPtr := ptr.Is(obj)

	var nonPtrObj interface{}
	var ptrObj interface{}

	if isPtr {
		ptrObj = obj
		// Let's convert to non ptr also!
		nonPtrObj = reflect.Indirect(reflect.ValueOf(obj)).Interface()
		//nonPtrObj = reflect.ValueOf(obj).Elem()

	} else {
		nonPtrObj = obj
	}

	// For non Ptr
	v := reflect.ValueOf(nonPtrObj)
	t := reflect.TypeOf(nonPtrObj)
	i := reflect.Indirect(v)

	// if it's pointer, then try getting the value of the structure and also save it!

	h := &Helper{
		TagName: DefaultTagName,

		PtrObj: ptrObj,

		// Non Ptr
		NonPtrObj:       nonPtrObj,
		valueOf:         v,
		typeOf:          t,
		indirectValueOf: i,

		// Other
		isPtr: isPtr,
	}

	if isPtr {
		h.ptrValueOf = reflect.ValueOf(ptrObj)
		h.ptrTypeOf = reflect.TypeOf(ptrObj)
		h.ptrIndirectValueOf = reflect.Indirect(h.ptrValueOf)
	}

	isStr := h.isStruct()
	h.isStr = isStr

	isPlainStr := h.isPlainStruct()
	h.isPlainStr = isPlainStr

	return h
}
