package main

import (
	"log"
	"time"
)

func main() {

	for {
		time.Sleep(time.Second)
		log.Println("hello world")
	}
}
