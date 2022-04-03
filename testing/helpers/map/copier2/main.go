package main

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/Map"
	"log"
	"reflect"
)

type Terminal struct {
	ID string
	//Name   string
	Name *string
	//UserID *uuid.UUID
	UserID uuid.UUID
}

func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
		isPtr = true
	}
	return reflectType, isPtr
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func main() {
	data := make(map[string]interface{})
	data["ID"] = "123456"
	data["Name"] = "Octavian"
	data["UserID"] = "9c6668c3-3605-424b-872c-c99c6140e113"
	//data["UserID"] = "9c6668c3"
	//data["UserID"] = ""

	var v1 *string
	var v3 interface{}
	v2 := "hello"
	//v1 = &v2
	//log.Println(v1)

	// in cazul in care type-ul nu a fost initializat/creat noi trebuie sa-l cream!
	// logic ar fi ca noi simple type sa cream apoi sa ii facem pointer sau sa chem de la acela pointer-ul, prin urmare o sa-i
	// setam pointer-ul celuilalt!

	// in caz ca e pointer eu scot element cu Elem!
	v1Type, isPtr := indirectType(reflect.TypeOf(v1))
	log.Println("is ptr", isPtr)
	// in cazul in care e pointer noi o sa-i cerem indirect

	var v1Ref reflect.Value
	if isPtr {
		// we should create the value
		v1Ref = indirect(reflect.New(v1Type))
	} else {
		//v1Ref = indirect(reflect.ValueOf(v1))
		v1Ref = indirect(reflect.New(v1Type))
	}
	log.Println("can set", v1Ref.CanSet())
	v1Ref.SetString("hello world")
	v3 = v1Ref.Addr()
	//newVal := reflect.New(v1Type)
	//v1Ref.Set(newVal)
	//
	//v1Ref.SetString("heeyoo")

	//log.Println("v1 val", *v1)
	log.Println("v1Ref val", v1Ref)
	log.Println("v1Ref val Addr", v1Ref.Addr())
	log.Println("v3", v3)
	log.Println("v1 type", v1Type)
	log.Println("v2", reflect.ValueOf(v2))
	log.Println("v2 type", reflect.TypeOf(v2))

	return

	newMap, _err := Map.ConvertMapValuesBasedOnModel(data, &Terminal{}, nil)

	if _err != nil {
		log.Println("error -> ", _err.Error())
	}
	log.Println(newMap)
	v := newMap["Name"].(**string)
	log.Println(v)
}
