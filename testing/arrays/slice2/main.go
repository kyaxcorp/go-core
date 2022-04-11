package main

import (
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"log"
	"reflect"
)

type Color struct {
	Name string
	//Category *ColorCategory
	Category *ColorCategory
}

type ColorCategory struct {
	CatName string
}

type str1 struct {
	Color *Color
}

func main() {
	//cat := &ColorCategory{CatName: "shiny"}

	//c := Color{Name: "red", Category: cat}
	c := &Color{Name: "red", Category: nil}
	v1 := str1{
		Color: c,
	}

	//v := _struct.GetNestedFieldReflectValue(reflect.ValueOf(v1), "Color.Name")
	//v, _err := _struct.GetNestedFieldReflectValue(reflect.ValueOf(v1), "Color.Category.CatName")
	v, _err := _struct.GetNestedFieldReflectValue(reflect.ValueOf(v1), "Color.Name")
	log.Println(v, _err)
}
