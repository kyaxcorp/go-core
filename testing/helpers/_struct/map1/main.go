package main

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_struct"
	"log"
)

type InputData struct {
	Name string
	Age  int
}

func main() {
	inputData := &InputData{
		Name: "Octavian",
		Age:  28,
	}
	inputData2 := InputData{
		Name: "Octavian",
		Age:  28,
	}

	log.Println(_struct.New(inputData).Map())
	log.Println(_struct.New(inputData2).Map())
}
