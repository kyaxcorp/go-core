package _struct

import (
	"github.com/google/uuid"
	"reflect"
)

// GetFieldValue -> return a structure field value
func GetFieldValue(obj interface{}, fieldName string) interface{} {
	return New(obj).GetFieldValue(fieldName)
}

func (h *Helper) GetFieldValue(fieldName string) interface{} {
	//if !IsPlainStruct(obj) {
	//	panic(InvalidObjType)
	//}
	f := h.indirectValueOf.FieldByName(fieldName)
	//f.CanInterface()
	return f.Interface()
}

func SetAny(obj interface{}, fieldName string, val interface{}) bool {
	return New(obj).SetAny(fieldName, val)
}

func (h *Helper) SetAny(fieldName string, val interface{}) bool {
	return h.SetInterface(fieldName, val)
}

func SetInterface(obj interface{}, fieldName string, val interface{}) bool {
	return New(obj).SetInterface(fieldName, val)
}

func (h *Helper) SetInterface(fieldName string, val interface{}) bool {
	//if !IsPlainStruct(obj) {
	//	panic(InvalidObjType)
	//}

	//if h.ptrIndirectValueOf.IsNil() {
	//	return false
	//}

	f := h.ptrIndirectValueOf.FieldByName(fieldName)
	v := reflect.ValueOf(val)

	objFieldType := f.Type().String()
	valFieldType := v.Type().String()

	if f.CanSet() {
		if objFieldType == "uuid.UUID" && valFieldType == "*uuid.UUID" {
			realVal := val.(*uuid.UUID)
			v = reflect.ValueOf(*realVal)
			f.Set(v)
			return true
		} else if objFieldType == "*uuid.UUID" && valFieldType == "uuid.UUID" {
			realVal := val.(uuid.UUID)
			v = reflect.ValueOf(&realVal)
			f.Set(v)
			return true
		}

		f.Set(v)
		return true
	}
	return false
}

func SetUUID(obj interface{}, fieldName string, val uuid.UUID) bool {
	return New(obj).SetUUID(fieldName, val)
}

func (h *Helper) SetUUID(fieldName string, val uuid.UUID) bool {
	//if !IsPlainStruct(obj) {
	//	panic(InvalidObjType)
	//}

	v := reflect.ValueOf(val)
	f := h.indirectValueOf.FieldByName(fieldName)

	if f.CanSet() {
		//f.SetBytes(val)
		// TODO: how to set bytes?!
		//f.SetBytes(val)
		f.Set(v)
		return true
	}
	return false
}

func SetBool(obj interface{}, fieldName string, val bool) bool {
	return New(obj).SetBool(fieldName, val)
}

func (h *Helper) SetBool(fieldName string, val bool) bool {
	//if !IsPlainStruct(obj) {
	//	panic(InvalidObjType)
	//}
	f := h.indirectValueOf.FieldByName(fieldName)
	if f.CanSet() {
		f.SetBool(val)
		return true
	}
	return false
}

func (h *Helper) Bool(fieldName string, val bool) bool {
	return SetBool(h.NonPtrObj, fieldName, val)
}

func GetPointerStructValue(obj interface{}) interface{} {
	return New(obj).GetPointerStructValue()
}

func (h *Helper) GetPointerStructValue() interface{} {
	//if !IsPointer(obj) || !IsStruct(obj) {
	// TODO: check also if it's a structure!
	if !IsPointer(h.NonPtrObj) {
		panic(InvalidPointerObjType)
	}
	return h.indirectValueOf.Interface()
}
