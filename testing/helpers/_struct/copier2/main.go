package main

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"log"
)

type Terminal struct {
	ID string
}

type Details struct {
	ID uuid.UUID
}

func copy(v1 interface{}, v2 interface{}) {
	_err := copier.Copy(v2, v1)
	if _err != nil {
		log.Println("failed to copy!!! " + _err.Error())
	}
}

func main() {
	v1 := &Terminal{
		ID: "3843ec50-a346-4e56-8008-9725aa925398",
	}
	v2 := &Details{}

	copy(v1, v2)
	log.Println(v2)
}
