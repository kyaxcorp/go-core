package main

import (
	"log"
	"reflect"
)

type Human struct {
	Name string
	Age  int
}

type TheCopier1 struct {
	From interface{}
	To   interface{}
}

type TheCopier2 struct {
	From *interface{}
	To   *interface{}
}

func main() {
	j := Human{
		Name: "Octavian",
		Age:  28,
	}
	j2 := &Human{
		Name: "Octavian",
		Age:  28,
	}
	j2pointer := &j2

	test := reflect.ValueOf(&j).Elem()
	test2 := reflect.Indirect(reflect.ValueOf(&j))
	test3 := reflect.Indirect(reflect.ValueOf(&j))
	test4 := reflect.Indirect(reflect.ValueOf(j))

	test5 := reflect.Indirect(reflect.ValueOf(&j2))
	test6 := reflect.Indirect(reflect.ValueOf(j2))

	var test7 interface{}
	test7 = reflect.Indirect(reflect.ValueOf(&j2)).Interface()
	test8 := reflect.ValueOf(&test7)
	test9 := reflect.ValueOf(&test7).Elem()
	test10 := reflect.ValueOf(&test7).Elem().Interface()

	log.Println("j 1")
	log.Println("j", j)
	log.Println("&j", &j)
	log.Println("j ValueOf", reflect.ValueOf(j))
	log.Println("&j ValueOf", reflect.ValueOf(&j))
	log.Println("j TypeOf", reflect.TypeOf(j))
	log.Println("&j TypeOf", reflect.TypeOf(&j))
	log.Println("j Indirect", reflect.Indirect(reflect.ValueOf(j)))
	log.Println("&j Indirect", reflect.Indirect(reflect.ValueOf(&j)))
	log.Println("j can addr", reflect.ValueOf(j).CanAddr())
	log.Println("&j can addr", reflect.ValueOf(&j).CanAddr())
	log.Println("test", test.CanAddr())
	log.Println("test2", test2.CanAddr())

	log.Println("test3", test3.CanAddr(), test3.Interface())
	log.Println("test4", test4.CanAddr(), test4.Interface())
	log.Println("test5", test5.CanAddr(), test5.Interface())
	log.Println("test6", test6.CanAddr(), test6.Interface())
	log.Println("test7", test6.CanAddr(), test6.Interface())
	log.Println("test8", test8.CanAddr(), test8.Interface())
	log.Println("test9", test9.CanAddr(), test9.Interface())
	log.Println("test10", test10)

	log.Println("j2")
	log.Println("j2", j2)
	log.Println("&j2", &j2)
	log.Println("j2pointer", j2pointer)
	log.Println("&j2pointer", &j2pointer)
	log.Println("j2 can addr", reflect.ValueOf(j2).CanAddr())
	log.Println("&j2 can addr", reflect.ValueOf(&j2).CanAddr())

	//reflect.Indirect()
	//reflect.TypeOf()
	//reflect.ValueOf()
}
