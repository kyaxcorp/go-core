package main

import (
	"github.com/KyaXTeam/go-core/core/helpers/slice"
	"log"
)

type Vasea struct {
	name string
}

func main() {
	// Here we define the slice
	var jordan []*Vasea

	if jordan == nil {
		log.Println("Jordan is nil!!!!")
	}
	// Here we just create an empty slice...
	jordan = []*Vasea{}
	// We check again if it's empty...
	if jordan == nil {
		log.Println("Jordan (2) is nil!!!!")
	}
	log.Println(jordan)

	jordan = append(jordan, &Vasea{name: "hey"})
	jordan = append(jordan, &Vasea{name: "hey"})
	//jordan := []*Vasea{}

	for key, val := range jordan {
		log.Println(key, val)
	}

	if status, _ := slice.IndexExists(jordan, 0); status {
		log.Println("yess")
	}
	if status, _ := slice.IndexExists(jordan, 10); status {
		log.Println("yess")
	}

	log.Println(jordan[0])
	log.Println(jordan)
}
