package _context

import (
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"log"
	"time"
)

func Start() {
	go func() {
		for {
			select {
			case <-_context.GetDefaultContext().Done():
				log.Println("Finished 1!!!")
			default:
				log.Println("hello")
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			select {
			case <-_context.GetDefaultContext().Done():
				log.Println("Finished 2!!!")
			}
			time.Sleep(time.Second)
		}
	}()

	time.AfterFunc(time.Second*3, func() {
		_context.Cancel()
	})

	for {
		time.Sleep(time.Second)
	}
}
