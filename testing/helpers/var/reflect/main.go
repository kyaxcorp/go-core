package main

import (
	"github.com/google/uuid"
	"log"
	"reflect"
)

type TestModel struct {
	UpdateBy  uuid.UUID
	UpdateBy2 *uuid.UUID
}

func main() {

	obj := &TestModel{}
	//var val interface{}

	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName("UpdateBy")
	f2 := reflect.Indirect(r).FieldByName("UpdateBy2")
	//v := reflect.ValueOf(val)
	log.Println("type", f.Type().String())
	log.Println("type 2", f2.Type().String())

}
