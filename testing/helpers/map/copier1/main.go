package main

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/Map"
	"github.com/google/uuid"
	"log"
)

type Terminal struct {
	ID     string
	Name   string
	UserID uuid.UUID
}

func main() {
	data := make(map[string]interface{})
	data["ID"] = "123456"
	data["Name"] = "Octavian"
	data["UserID"] = "9c6668c3-3605-424b-872c-c99c6140e113"
	//data["UserID"] = "9c6668c3"
	//data["UserID"] = ""

	newMap, _err := Map.ConvertMapValuesBasedOnModel(data, &Terminal{}, nil)

	if _err != nil {
		log.Println("error -> ", _err.Error())
	}
	log.Println(newMap)
}
