package main

import (
	"log"
	"reflect"
)

type Car struct {
	Name string
}

func CreateArray(t reflect.Type, length int) reflect.Value {
	var arrayType reflect.Type
	arrayType = reflect.ArrayOf(length, t)
	return reflect.Zero(arrayType)
}

func main() {
	var v1 interface{}
	var v2 interface{}

	v1 = &Car{}
	log.Println(reflect.ValueOf(v1))
	log.Println(reflect.TypeOf(v1))
	log.Println(reflect.Indirect(reflect.ValueOf(v1)))

	//typeOf := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(v1)).Interface())
	typeOf := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(v1)).Interface())
	//log.Println(typeOf)

	slice := reflect.SliceOf(typeOf)
	v2 = slice
	log.Println(v2)

}
