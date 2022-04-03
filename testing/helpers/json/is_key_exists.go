package main

import (
	"github.com/KyaXTeam/go-core/core/helpers/json"
	"log"
)

func main() {
	myJson := `
		{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "name2": {"first": "", "last": "Anderson"},
	  "name3": {"last": "Anderson"},
      "name4": {"first": 0, "last": "Anderson"},
	  "age":37,
	  "children": ["Sara","Alex","Jack"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "James", "last": "Murphy"},
		{"first": "Roger", "last": "Craig"}
	  ]
	}
	`
	log.Println(json.IsKeyExists(myJson, "name.first"))
	log.Println(json.IsKeyExists(myJson, "name2.first"))
	log.Println(json.IsKeyExists(myJson, "name3.first"))
	log.Println(json.IsKeyExists(myJson, "name4.first"))
}
