package main

import "log"

var test = make(map[string]int)

func main() {
	test["vasea"] = 2

	log.Println(test)

}
