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

	if exists, _ := slice.IndexExists(jordan, 0); exists {
		log.Println("exists!")
	}

	// Here we just create an empty slice...
	jordan = append(jordan, &Vasea{})
	log.Println(jordan)

	if exists, _ := slice.IndexExists(jordan, 0); exists {
		log.Println("exists!")
	}
}
