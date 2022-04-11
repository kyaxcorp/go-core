package main

import (
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"log"
	"reflect"
)

type Color struct {
	Name     string
	Category ColorCategory
}

type ColorCategory struct {
	CatName string
}

type str1 struct {
	Color Color
}

func main() {
	cat := ColorCategory{CatName: "shiny"}

	c := Color{Name: "red", Category: cat}
	v1 := str1{
		Color: c,
	}

	//v := _struct.GetNestedFieldReflectValue(reflect.ValueOf(v1), "Color.Name")
	v := _struct.GetNestedFieldReflectValue(reflect.ValueOf(v1), "Color.Category.CatName")
	log.Println(v)
}
