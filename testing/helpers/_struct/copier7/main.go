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
	ID *uuid.UUID
}

type TheCopier struct {
	From interface{}
	To   interface{}
}

func (c *TheCopier) copy() {
	_err := copier.Copy(&c.To, c.From)
	if _err != nil {
		log.Println("failed to copy!!! " + _err.Error())
	}
}

func main() {
	id := "3843ec50-a346-4e56-8008-9725aa925398"
	from := Terminal{
		ID: id,
	}
	to := Details{}

	c := &TheCopier{
		From: from,
		To:   to,
	}
	c.copy()

	log.Println(c.To)
}
