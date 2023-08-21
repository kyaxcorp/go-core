package main

import (
	"log"
	"reflect"
)

type Test struct {
	MyField *Test2
}

type Test2 struct {
	Name string
}

func main() {

	obj := &Test{}
	log.Println(obj)

	log.Println(reflect.TypeOf(obj))
	log.Println(reflect.TypeOf(obj).Elem())
	//log.Println(reflect.TypeOf(obj).Elem().Elem()) => IT GIVES ERROR

	if reflect.TypeOf(obj).Kind() == reflect.Pointer {
		log.Println("YES IT'S A POINTER")
	}

	if reflect.TypeOf(obj).Elem().Kind() == reflect.Struct {
		log.Println("YES IT'S A STRUCT")
	}

	log.Println(reflect.New(reflect.TypeOf(obj)))
	log.Println(reflect.New(reflect.TypeOf(obj).Elem()))

	log.Println()
}
