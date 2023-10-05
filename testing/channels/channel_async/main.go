package main

import (
	"github.com/google/uuid"
	"log"
	"time"
)

func main() {

	myChannel := make(chan uuid.UUID)

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
		log.Println("pushing", i)
		id, _ := uuid.NewRandom()
		myChannel <- id
		log.Println("pushed", i)
	}

	log.Println("finished pushing data to channel")

	time.Sleep(time.Second * 10)
}
