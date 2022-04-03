package events

import (
	"github.com/KyaXTeam/go-core/core/helpers/event"
	"log"
	"time"
)

func main() {
	event.Listen("record.new", &event.ListenOptions{
		Name: "on_create",
		Callback: func(dispatchData *event.DispatchData) {
			log.Println("1111 i have received an event!!!")
			log.Println(dispatchData)
		},
		Async:             true,
		ExecutionPriority: 10,
		OnEventRegisterFinish: func(id string) {
			log.Println("registration success", id)
		},
	})

	event.Listen("record.new", &event.ListenOptions{
		Name: "on_create222",
		Callback: func(dispatchData *event.DispatchData) {
			log.Println("2222i have received an event!!!")
			log.Println(dispatchData)
		},
		Async:             true,
		ExecutionPriority: 9,
		OnEventRegisterFinish: func(id string) {
			log.Println("registration success", id)
		},
	})

	log.Println("listening started!")

	go func() {
		for {
			log.Println("dispatching...")
			myData := "hello world"
			event.Dispatch("record.new", &event.DispatchData{
				Id:   32323,
				Data: myData,
			})
			time.Sleep(time.Second)
		}
	}()

	for {
		time.Sleep(time.Second)
		//log.Println(len(mymap))
	}
}
