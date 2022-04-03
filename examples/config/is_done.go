package config

import (
	"github.com/KyaXTeam/go-core/core/helpers/_context"
	"log"
	"time"
)

func main() {
	ctx := _context.WithCancel(_context.GetDefaultContext())

	go func() {
		time.Sleep(time.Second * 5)
		ctx.Cancel()
	}()

	go func() {
		for {
			log.Println("sleeping...")
			time.Sleep(time.Second)
			if ctx.IsDone() {
				log.Println("2 - IS DONE!!!")
			} else {
				log.Println("2 - NO!")
			}
		}
	}()

	for {
		log.Println("sleeping...")
		time.Sleep(time.Second)
		if ctx.IsDone() {
			log.Println("1 - IS DONE!!!")
		} else {
			log.Println("1 - NO!")
		}
	}
}
