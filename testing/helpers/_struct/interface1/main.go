package main

import (
	"log"
	"reflect"
)

type TheCopier1 struct {
	From interface{}
	To   interface{}
}

type TheCopier2 struct {
	From *interface{}
	To   *interface{}
}

func main() {
	// Test 1
	log.Println("test 1")
	log.Println(TheCopier1{})
	log.Println(TheCopier2{})

	// Test 2
	log.Println("test 2")
	log.Println(TheCopier1{
		From: nil,
		To:   nil,
	})
	log.Println(TheCopier2{
		From: nil,
		To:   nil,
	})
	// Test 3
	log.Println("test 3")
	log.Println(TheCopier1{
		From: "hello world",
		To:   33333,
	})
	t3v1 := "hello world"
	t3v2 := 33333
	log.Println(TheCopier2{
		From: &t3v1,
		To:   &t3v2,
	})

	//reflect.TypeOf()
	//reflect.ValueOf()
}
