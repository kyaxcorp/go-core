package main

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/helpers/json"
	"log"
)

type Terminal struct {
	ID     string
	Name   string
	UserID uuid.UUID
}

func main() {
	data := make(map[string]interface{})

	//userID, _ := uuid.Parse()
	//data["ID"] = "123456"
	//data["Name"] = "Octavian"
	//data["UserID"] = userID

	data["ID"] = "123456"
	data["Name"] = "Octavian"
	data["UserID"] = "9c6668c3-3605-424b-872c-c99c6140e113"

	_json, _ := json.Encode(data)

	log.Println(data)

	t := &Terminal{}

	json.Decode(_json, t)
	//_err := mapstructure.Decode(data, t)
	//
	////_err := copier.Copy(t, data)
	//if _err != nil {
	//	log.Println("error", _err)
	//}
	log.Println(t)

	//Map.ConvertMapValuesBasedOnModel()
}
