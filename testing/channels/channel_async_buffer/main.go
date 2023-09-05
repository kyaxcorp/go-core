package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

func main() {

	myChannel := make(chan uuid.UUID, 2)

	go func() {
		for {
			select {
			case data := <-myChannel:
				log.Println(data, "sleeping... 3 sec")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		id, _ := uuid.NewRandom()
		log.Println("pushing", id)
		//go func() {
		myChannel <- id
		//}()
		log.Println("pushed", id)
	}

	log.Println("finished pushing data to channel")

	time.Sleep(time.Second * 10)
}
