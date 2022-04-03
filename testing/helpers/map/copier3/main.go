package main

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/Map"
	"github.com/google/uuid"
	"log"
)

type Terminal struct {
	ID string
	//Name   string
	Name *string
	//UserID *uuid.UUID
	UserID uuid.UUID
}

func main() {
	data := make(map[string]interface{})
	data["ID"] = "123456"
	//data["Name"] = "Octavian"
	data["Name"] = nil
	//data["UserID"] = "9c6668c3-3605-424b-872c-c99c6140e113"
	data["UserID"] = nil

	newMap, _err := Map.ConvertMapValuesBasedOnModel(data, &Terminal{}, nil)

	if _err != nil {
		log.Println("error -> ", _err.Error())
	}
	log.Println(newMap)
	//v := newMap["Name"].(*string)
	//log.Println(v)
}
