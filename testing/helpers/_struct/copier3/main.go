package main

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"log"
)

type Terminal struct {
	ID *string
}

type Details struct {
	ID uuid.UUID
}

func copy(from interface{}, to interface{}) {
	_err := copier.Copy(to, from)
	if _err != nil {
		log.Println("failed to copy!!! " + _err.Error())
	}
}

func main() {
	id := "3843ec50-a346-4e56-8008-9725aa925398"
	from := &Terminal{
		ID: &id,
	}
	to := &Details{}

	log.Println(from, *from.ID)
	copy(from, to)
	log.Println(to)
}
