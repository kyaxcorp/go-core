package _struct

import "reflect"

type Helper struct {
	// this is the plain structure!
	NonPtrObj interface{}
	// this is the pointer to structure!
	PtrObj interface{}

	// the input that we received it is a pointer of a structure or not
	isPtr bool
	// the input that we received it is a structure or not...
	isStr bool
	// the input that we received is plain structure or not
	isPlainStr bool

	// These are the values and types of the NonPtrObject
	valueOf         reflect.Value
	typeOf          reflect.Type
	indirectValueOf reflect.Value

	// Ptr
	ptrValueOf         reflect.Value
	ptrTypeOf          reflect.Type
	ptrIndirectValueOf reflect.Value

	TagName string
}
