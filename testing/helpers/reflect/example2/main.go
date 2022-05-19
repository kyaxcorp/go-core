package main

import (
	"log"
	"reflect"
)

type Car struct {
	Name *string
}

func main() {
	vasea := "hellow wolrd"
	v1 := &Car{Name: &vasea}

	var myVar interface{}
	myVar = v1.Name
	refType := reflect.TypeOf(myVar)
	refTypeNative := refType
	refKind := refType.Kind()
	refVal := reflect.ValueOf(myVar)
	refValNative := refVal
	refIsPtr := false

	if refKind == reflect.Ptr {
		refIsPtr = true
	}

	if refIsPtr {
		refTypeNative = refType.Elem()
		if refVal.IsZero() {
			// if it's zero, we should create an empty zero value
			refValNative = reflect.Zero(refTypeNative)
		} else {
			// We should take the real indirect type value
			refValNative = reflect.Indirect(refVal)
		}
	}

	log.Println(refValNative.Interface())

	log.Println(reflect.TypeOf(v1.Name))
	log.Println(reflect.TypeOf(v1.Name).Elem())
	log.Println(reflect.TypeOf(v1.Name).Kind())
	log.Println(reflect.ValueOf(v1.Name).Interface())
	log.Println(reflect.ValueOf(v1.Name).IsZero())
	//log.Println(reflect.Indirect(reflect.ValueOf(v1.Name)).IsZero())
}
