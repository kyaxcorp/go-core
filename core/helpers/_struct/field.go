package _struct

import (
	"reflect"
	"strings"
)

// TODO: move or rename this func
func GetExportableFieldsByJSONTag(obj interface{}) []string {
	val := reflect.ValueOf(obj)
	var exportable []string

	for i := 0; i < val.Type().NumField(); i++ {
		t := val.Type().Field(i)
		fieldName := t.Name

		switch jsonTag := t.Tag.Get("json"); jsonTag {
		case "-":
		case "":
			// fmt.Println(fieldName)
			exportable = append(exportable, fieldName)
		default:
			parts := strings.Split(jsonTag, ",")
			name := parts[0]
			if name == "" {
				name = fieldName
			}
			exportable = append(exportable, name)
			//fmt.Println(name)
		}
	}

	return exportable
}

func FieldExists(obj interface{}, fieldName string) bool {
	return New(obj).FieldExists(fieldName)
}

func (h *Helper) FieldExists(fieldName string) bool {
	f := h.indirectValueOf.FieldByName(fieldName)
	if f == (reflect.Value{}) {
		return false
	}
	return true
}

func GetFieldType(obj interface{}, fieldName string) string {
	return New(obj).GetFieldType(fieldName)
}

func (h *Helper) GetFieldType(fieldName string) string {
	f := h.indirectValueOf.FieldByName(fieldName)
	if f == (reflect.Value{}) {
		return ""
	}
	return f.Type().Kind().String()
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
		isPtr = true
	}
	return reflectType, isPtr
}

func (h *Helper) GetFieldReflectType(fieldName string) reflect.Type {
	// TODO Check each field by type and based on that use .Elem()

	f := h.indirectValueOf.FieldByName(fieldName)
	if f == (reflect.Value{}) {
		return nil
	}

	return f.Type()
	//return indirectType(reflect.TypeOf(f.Interface()))
	//return reflect.TypeOf(f.Interface())
}

func GetFieldTypeName(obj interface{}, fieldName string) string {
	return New(obj).GetFieldTypeName(fieldName)
}

func (h *Helper) GetFieldTypeName(fieldName string) string {
	f := h.indirectValueOf.FieldByName(fieldName)
	if f == (reflect.Value{}) {
		return ""
	}
	return f.Type().Name()
}
