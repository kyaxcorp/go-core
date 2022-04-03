package context

import (
	"context"
	"log"
	"time"
)

var GlobalContext context.Context

func main() {
	ctx, cancel := context.WithCancel(GlobalContext)
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	go func() {
		for {
			select {
			case v := <-ctx.Done():
				log.Println(v)
				time.Sleep(time.Second)
			default:
				log.Println("nothing intersting...")
				time.Sleep(time.Second)
			}
		}
	}()

}
